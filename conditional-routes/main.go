package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const messageKey = "message"

func main() {
	enableAdmin := flag.Bool("admin", false, "Enable admin routes")
	flag.Parse()

	enableMetrics := os.Getenv("ENABLE_METRICS") == "true"

	r := gin.Default()

	// Public routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			messageKey: "Welcome to Gin",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})

	// Conditional admin routes
	if *enableAdmin {
		admin := r.Group("/admin")
		{
			admin.GET("/users", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					messageKey: "Admin users endpoint",
				})
			})

			admin.GET("/settings", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					messageKey: "Admin settings endpoint",
				})
			})
		}
	}

	// Conditional metrics route
	if enableMetrics {
		r.GET("/metrics", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"requests_total": 1234,
				"uptime":         "24h",
			})
		})
	}

	log.Fatal(r.Run(":8080"))
}
