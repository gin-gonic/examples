package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yosssi/ace"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ace-example", func(c *gin.Context) {
		tpl, err := ace.Load("./ace-example", "", nil)
		if err != nil {
			panic(err)
		}

		err = tpl.Execute(c.Writer, gin.H{"Title": "Gin"})
		if err != nil {
			panic(err)
		}
	})

	return router
}

func main() {
	r := setupRouter()

	r.Run(":3333")
}
