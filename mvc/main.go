package main

import (
	"mvc-example/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func main() {
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	user := r.Group("/user")
	router.UserRouter(user)
	r.Static("/user/public", "./public")
	r.LoadHTMLGlob("views/*")

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://locolhost"}
	r.Use(cors.New(config))

	r.Run()
}
