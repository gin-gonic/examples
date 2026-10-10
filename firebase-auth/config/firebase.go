package config

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

const (
	ErrorFirebaseInit = "error initializing Firebase app"
	ErrorAuthClient = "error getting Auth client"
)
func InitializeFirebase() (*auth.Client, error){
	//add cred file downloaded from project settings -> service accounts
	opt := option.WithCredentialsFile("/Users/vinayak/Documents/oss/examples/firebase-auth/config/cred.json")
	
	//init new app on firebasea
	app, err := firebase.NewApp(context.Background(), nil , opt)
	if err != nil {
		log.Fatalf(ErrorFirebaseInit+": %v\n",err)
		return nil, fmt.Errorf(ErrorFirebaseInit)
	}

	//init new auth app client
	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf(ErrorAuthClient+": %v\n", err)
		return nil, err
	}

	return authClient, nil

}