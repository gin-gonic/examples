package main

import (
	"log"

	"github.com/gin-gonic/examples/firebase-auth/config"
	"github.com/gin-gonic/examples/firebase-auth/routes"
	"github.com/gin-gonic/gin"
)

const (
	ErrStartServer = "error starting server"
	PORT = "8080"
)
func main(){

	//init firebase
	firebaseAuth, err := config.InitializeFirebase()
	if err != nil {
		log.Fatalf(config.ErrorFirebaseInit+": %v",err)
	}

	//set up router
	router := gin.Default()

	//register routes
	routes.RegisterRoutes(router, firebaseAuth)

	//start server
	if err := router.Run(":"+PORT); err != nil {
		log.Fatalf("err starting server ")
	}

}