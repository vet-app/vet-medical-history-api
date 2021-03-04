package helpers

import (
	"cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"
)

var FirebaseApp *firebase.App

func FirebaseConnection() {
	sdkFilePath := os.Getenv("FIREBASE_CONFIG")

	opt := option.WithCredentialsFile(sdkFilePath)
	config := &firebase.Config{
		ProjectID:     "vet-app-ui",
		StorageBucket: os.Getenv("FIREBASE_BUCKET"),
	}

	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatal("On Firebase SDK connection: ", err)
	}

	log.Println("Connected successfully to Firebase Admin SDK")
	FirebaseApp = app
}

func storageConnection() *storage.BucketHandle {
	client, err := FirebaseApp.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}

	return bucket
}

func StoreFile(file multipart.File, bucketName, filename string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	wc := storageConnection().Object(bucketName + filename).NewWriter(ctx)

	if _, err = io.Copy(wc, file); err != nil {
		return err
	}

	if err := wc.Close(); err != nil {
		return err
	}

	return nil
}
