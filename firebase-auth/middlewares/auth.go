package middlewares

import (
	"context"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/examples/firebase-auth/utils"
	"github.com/gin-gonic/gin"
)


func AuthMiddleware(firebaseAuth *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			utils.RespondWithError(c, http.StatusUnauthorized, "Authorization header required")
			c.Abort()
			return 
		}

		decodeToken, err := firebaseAuth.VerifyIDToken(context.Background(), token)
		if err != nil {
			utils.RespondWithError(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		c.Set("user", decodeToken.UID)
		c.Next()

	}
}