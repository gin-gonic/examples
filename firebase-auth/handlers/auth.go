package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/examples/firebase-auth/utils"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

//turn on auth type on firebase console what ever type you wanna use for auth

type LoginRequest struct {
	Token string `json:"token" binding:"required"`
}

type SignupRequest struct {
	Email string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}


type GoogleSignupRequest struct {
	IDToken string `json:"id_token" binding:"required"`
}

//Signup handler
func SignupHandler(firebaseAuth *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SignupRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid request" )
			return 
		}

		existing, _ := firebaseAuth.GetUserByEmail(context.Background(), req.Email)
		if existing != nil {
			utils.RespondWithError(c, http.StatusNotAcceptable, "User already exisit with this email")
			return
		}

		params := (&auth.UserToCreate{}).
				Email(req.Email).
				Password(req.Password).
				DisplayName(req.Username)


		newUser, err := firebaseAuth.CreateUser(context.Background(), params)
		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create user")
			return
		}		

		c.JSON(http.StatusCreated, gin.H{
			"message":"User Created",
			"uid":newUser.UID,
		})
	}
}

func SigninWithGoogleHandler(firebaseAuth *auth.Client)gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GoogleSignupRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid Request")
			return
		}

		//verify google 
		gUser, err := firebaseAuth.VerifyIDToken(context.Background(), req.IDToken)
		if err!= nil {
			utils.RespondWithError(c, http.StatusBadRequest,"Failed to verify google id token")
			return
		}

		c.Set("user", gUser.UID)
		c.JSON(http.StatusAccepted, gin.H{
			"message":"Login successful via google",
			"uid": gUser.UID,
		})		

	}

}

type EmailLoginRequest struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type FirebaseAuthResponse struct {
    IdToken string `json:"idToken"`
    Error   struct {
        Message string `json:"message"`
    } `json:"error"`
}

// Firebase project-specific endpoint for REST API
const firebaseAuthURL = "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=YOUR_FIREBASE_WEB_API_KEY"

// EmailLoginHandler handles user login with email and password
func EmailLoginHandler(firebaseAuth *auth.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
		fmt.Println("step 0")

		var req EmailLoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid request" )
			return 
		}

		fmt.Println("step 1")

		_,err := firebaseAuth.GetUserByEmail(c, req.Email)
		if err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid Email" )
			return 
		}

        // Create the payload for Firebase REST API
        payload := map[string]string{
            "email":             req.Email,
            "password":          req.Password,
            "returnSecureToken": "true",
        }
        payloadBytes, _ := json.Marshal(payload)

		fmt.Println("step 2")

        // Make the request to Firebase REST API
        resp, err := http.Post(firebaseAuthURL, "application/json", bytes.NewBuffer(payloadBytes))
        if err != nil || resp.StatusCode != http.StatusOK {
            utils.RespondWithError(c, http.StatusUnauthorized, "Invalid email or password")
            return
        }
        defer resp.Body.Close()

		fmt.Println("step 3")

        // Parse the response
        body, _ := ioutil.ReadAll(resp.Body)
        var authResp FirebaseAuthResponse
        if err := json.Unmarshal(body, &authResp); err != nil {
            utils.RespondWithError(c, http.StatusInternalServerError, "Failed to parse response from Firebase")
            return
        }
		fmt.Println("step 4")

        if authResp.IdToken == "" {
            utils.RespondWithError(c, http.StatusUnauthorized, authResp.Error.Message)
            return
        }
		fmt.Println("step 5")

        c.JSON(http.StatusOK, gin.H{
            "message": "Login successful",
            "token":   authResp.IdToken,
        })
    }
}

func LoginHandler(firebaseAuth *auth.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req LoginRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid Request")
			return
		}

		token, err := firebaseAuth.VerifyIDToken(context.Background(), req.Token)
		if err != nil {
			utils.RespondWithError(ctx, http.StatusUnauthorized, "Invalid Token")
			return
		}


		ctx.JSON(http.StatusAccepted, gin.H{
			"message":"Login successfull",
			"uid": token.UID,
		})

	}
}

func ProfileHandler(c *gin.Context){
	user, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message":"profile not found",
		})
	}
	c.JSON(http.StatusFound, gin.H{
		"message":"profile found",
		"user":user,
	})
}