package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CookieTool() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get cookie
		if cookie, err := c.Cookie("label"); err == nil {
			if cookie == "ok" {
				c.Next()
				return
			}
		}

		// Cookie verification failed
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden with no cookie"})
		c.Abort()
	}
}

func main() {
	route := gin.Default()

	route.GET("/login", func(c *gin.Context) {
		// Set cookie {"label": "ok" }, maxAge 30 seconds.
		c.SetCookie("label", "ok", 30, "/", "localhost", false, true)
		c.String(200, "Login success!")
	})

	route.GET("/home", CookieTool(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "Your home page"})
	})

	route.Run(":8080")
}
