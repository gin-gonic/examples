package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(RequestID())

	r.GET("/hello", func(c *gin.Context) {
		requestID, _ := c.Get("RequestID")

		c.JSON(http.StatusOK, gin.H{
			"message":    "Hello from Gin!",
			"request_id": requestID,
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
