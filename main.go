package main

import (
	"github.com/vet-app/vet-medical-history-api/pkg"
	"os"
)

var server = pkg.Server{}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("$PORT not set")
	}

	server.Initialize(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("SSL_MODE"),
	)

	server.Run(":" + port)
}
