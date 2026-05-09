package main

import (
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
)

var portPattern = regexp.MustCompile(`^[0-9]{1,5}$`)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	} else if !portPattern.MatchString(port) {
		log.Fatalf("invalid PORT value")
	}

	// Starts a new Gin instance with no middle-ware
	r := gin.New()

	// Define handlers
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Listen and serve on defined port
	log.Printf("Listening on port %s", port)
	_ = r.Run(":" + port)
}
