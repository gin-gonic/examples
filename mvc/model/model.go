package model

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"title": "Hello World",
	})
}
