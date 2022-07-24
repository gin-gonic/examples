package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	// serve static favicon file from a location relative to main.go directory
	//app.StaticFile("/favicon.ico", "./.assets/favicon.ico")
	app.StaticFile("/favicon.ico", "./favicon.ico")

	app.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello favicon.")
	})

	app.Run(":8080")
}
