package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Access-Control-Max-Age", "86400")
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		log.Println("Received a GET /ping request.")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"time":    time.Now().Format(time.RFC3339),
		})
	})
	r.POST("/data", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"action":  "Created",
			"message": "Data processed successfully via POST.",
		})
	})
	r.PUT("/data/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"action":  "Updated",
			"message": fmt.Sprintf("Successfully processed PUT for resource ID: %s", id),
		})
	})
	r.DELETE("/data/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"action":  "Deleted",
			"message": fmt.Sprintf("Successfully processed DELETE for resource ID: %s", id),
		})
	})
	log.Fatal(r.Run(":8080"))
}
