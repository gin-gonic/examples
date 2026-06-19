package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Constants for JSON response keys.
// This helps avoid repeating string literals across the codebase.
const (
	keyMessage  = "message"
	keyStatus   = "status"
	keyAction   = "action"
	statusValue = "success"
)

// CORSMiddleware enables cross-origin requests.
// This is useful when frontend and backend are running separately.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Allow requests from any origin.
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// Allowed HTTP methods.
		c.Writer.Header().
			Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")

		// Allowed headers.
		c.Writer.Header().
			Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")

		// Allow credentials (cookies, auth headers, etc.).
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight OPTIONS request.
		if c.Request.Method == "OPTIONS" {
			// Cache preflight response for 24 hours.
			c.Writer.Header().Set("Access-Control-Max-Age", "86400")

			// Return 204 No Content for preflight.
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// Pass control to the next middleware/handler.
		c.Next()
	}
}

func main() {
	// Create a Gin router with default middleware:
	// logger and recovery middleware.
	r := gin.Default()

	// Register custom CORS middleware.
	r.Use(CORSMiddleware())

	// GET endpoint for health check / connectivity test.
	r.GET("/ping", func(c *gin.Context) {
		// Log incoming request.
		log.Println("Received a GET /ping request.")

		// Return JSON response.
		c.JSON(http.StatusOK, gin.H{
			keyMessage: "pong",

			// Current server time in RFC3339 format.
			"time": time.Now().Format(time.RFC3339),
		})
	})

	// POST endpoint for processing data.
	r.POST("/data", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			keyStatus:  statusValue,
			keyAction:  "Created",
			keyMessage: "Data processed successfully via POST.",
		})
	})

	// PUT endpoint for updating a resource by ID.
	r.PUT("/data/:id", func(c *gin.Context) {
		// Extract ID from URL parameter.
		id := c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			keyStatus: statusValue,
			keyAction: "Updated",

			// Response message including resource ID.
			keyMessage: fmt.Sprintf(
				"Successfully processed PUT for resource ID: %s",
				id,
			),
		})
	})

	// DELETE endpoint for removing a resource by ID.
	r.DELETE("/data/:id", func(c *gin.Context) {
		// Extract ID from URL parameter.
		id := c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			keyStatus: statusValue,
			keyAction: "Deleted",

			// Response message including resource ID.
			keyMessage: fmt.Sprintf(
				"Successfully processed DELETE for resource ID: %s",
				id,
			),
		})
	})

	// Start HTTP server on port 8080.
	// log.Fatal ensures the app exits if the server fails to start.
	log.Fatal(r.Run(":8080"))
}
