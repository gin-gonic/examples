package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	enableAdmin := flag.Bool("admin", false, "Enable admin routes")
	flag.Parse()

	enableMetrics := os.Getenv("ENABLE_METRICS") == "true"

	r := gin.Default()

	// Public routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Gin",
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
					"message": "Admin users endpoint",
				})
			})

			admin.GET("/settings", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "Admin settings endpoint",
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

	r.Run(":8080")
}
