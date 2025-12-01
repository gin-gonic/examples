package main

import (
	"context"
	"crypto/rand"
	"embed"
	"encoding/base64"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var (
	clientID     = os.Getenv("GITHUB_CLIENT_ID")
	clientSecret = os.Getenv("GITHUB_CLIENT_SECRET")
	redirectURL  = "http://localhost:8080/callback"
)

var (
	// Cache for storing state
	stateCache = cache.New(10*time.Minute, 20*time.Minute)

	// OAuth2 configuration
	oauth2Config = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     github.Endpoint,
		Scopes:       []string{"user:email", "read:user"},
	}

	// GitHub user information API
	userInfoURL = "https://api.github.com/user"
)

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

type GitHubUser struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

//go:embed templates/*
var templatesFS embed.FS

func main() {
	r := gin.Default()

	// Set HTML templates
	r.SetHTMLTemplate(template.Must(template.ParseFS(templatesFS, "templates/*")))

	// Home route
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "GitHub OAuth Example",
		})
	})

	// Login route - Redirect to GitHub for authentication
	r.GET("/login", func(c *gin.Context) {
		state, err := generateRandomState()
		if err != nil {
			c.String(http.StatusInternalServerError, "Unable to generate state value")
			return
		}

		// Store state for later verification
		stateCache.Set(state, true, cache.DefaultExpiration)

		// Redirect to GitHub for authentication
		authURL := oauth2Config.AuthCodeURL(state)
		c.Redirect(http.StatusFound, authURL)
	})

	// Callback route - Handle GitHub authentication callback
	r.GET("/callback", func(c *gin.Context) {
		// Retrieve and verify state
		state := c.Query("state")
		if _, exists := stateCache.Get(state); !exists {
			c.String(http.StatusBadRequest, "Invalid state value")
			return
		}
		stateCache.Delete(state)

		// Retrieve code
		code := c.Query("code")
		if code == "" {
			c.String(http.StatusBadRequest, "Authorization code not provided")
			return
		}

		// Exchange code for access token
		token, err := oauth2Config.Exchange(context.Background(), code)
		if err != nil {
			c.String(
				http.StatusInternalServerError,
				"Unable to exchange access token: "+err.Error(),
			)
			return
		}

		// Use access token to get user information
		client := oauth2Config.Client(context.Background(), token)
		resp, err := client.Get(userInfoURL)
		if err != nil {
			c.String(
				http.StatusInternalServerError,
				"Unable to retrieve user information: "+err.Error(),
			)
			return
		}
		defer resp.Body.Close()

		var user GitHubUser
		if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
			c.String(
				http.StatusInternalServerError,
				"Unable to parse user information: "+err.Error(),
			)
			return
		}

		// Display user information on success page
		c.HTML(http.StatusOK, "success.html", gin.H{
			"title":      "Authentication Successful",
			"username":   user.Login,
			"name":       user.Name,
			"email":      user.Email,
			"avatar_url": user.AvatarURL,
		})
	})

	// Protected route - Requires authentication to access
	r.GET("/protected", func(c *gin.Context) {
		// In a real application, session or JWT token should be checked here
		// This is just a simplified example
		c.String(http.StatusOK, "This is a protected resource!")
	})

	// Start server
	log.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}
