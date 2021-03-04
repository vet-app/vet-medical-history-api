package helpers

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
	"os"
)

var FirebaseApp *firebase.App

func FirebaseConnection() {
	sdkFilePath := os.Getenv("FIREBASE_CONFIG")

	opt := option.WithCredentialsFile(sdkFilePath)
	config := &firebase.Config{ProjectID: "vet-app-ui"}

	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatal("On Firebase SDK connection: ", err)
	}

	log.Println("Connected successfully to Firebase Admin SDK")
	FirebaseApp = app
}