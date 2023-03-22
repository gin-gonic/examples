package main

import (
	"github.com/gin-gonic/examples/upload-file/chunk/model"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.POST("/chunkUploadFile", model.ChunkUploadFile)
	g.Run(":8080")
}
