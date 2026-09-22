package routes

import (
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/examples/firebase-auth/handlers"
	"github.com/gin-gonic/examples/firebase-auth/middlewares"
	"github.com/gin-gonic/gin"
)

//register

func RegisterRoutes(router *gin.Engine, firebaseAuth *auth.Client){

	//unprotected
	router.POST("/signup", handlers.SignupHandler(firebaseAuth))
	router.POST("/login", handlers.LoginHandler(firebaseAuth))
	router.POST("/email", handlers.EmailLoginHandler(firebaseAuth))

	//using google token(third party)
	router.POST("/google", handlers.SigninWithGoogleHandler(firebaseAuth))
	
	//protected routes
	protected := router.Group("/protected")
	protected.Use(middlewares.AuthMiddleware(firebaseAuth))
	{
		protected.GET("/profile", handlers.ProfileHandler)
	}

}