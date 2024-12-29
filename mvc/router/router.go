package router

import (
	"mvc-example/model"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func UserRouter(rg *gin.RouterGroup) {
	rg.GET("/", model.Index)
}
