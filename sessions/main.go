package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/login", "./static")

	router.POST("/login", func(c *gin.Context) {
		// Get user login info
		username := c.PostForm("username")
		password := c.PostForm("password")

		// Authentication
		expectedPassword, ok := users[username]
		if !ok || expectedPassword != password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect username or password"})
			return
		}

		// Create a session token
		sessionToken := fmt.Sprintf("session-token-%s", username)
		expireTime := time.Now().Add(30 * time.Second)

		// Set session map in server
		sessions[sessionToken] = session{
			username:   username,
			expireTime: expireTime,
		}

		// Finally set cookie for client
		c.SetCookie("session_token", sessionToken, 60, "/", "localhost", false, true)
		c.JSON(200, gin.H{"data": "Login success!"})
	})

	router.GET("/home", func(c *gin.Context) {
		// Get and verify session_token
		if sessionToken, err := c.Cookie("session_token"); err == nil {
			userSession, exists := sessions[sessionToken]

			// If the session token not present in session map, return unauthorized error
			if !exists {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "please login first"})
				return
			}

			// If the session token is expird, return unauthorized error
			if userSession.Expired() {
				delete(sessions, sessionToken)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "your auth info is expired"})
				return
			}

			// If the session token is valid, show the welcome message
			c.JSON(200, gin.H{"data": fmt.Sprintf("Welcome to your home page: %s", userSession.username)})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "can not get your auth info, please login first"})
		}
	})

	router.Run(":8080")
}
