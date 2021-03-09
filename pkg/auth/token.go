package auth

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/vet-app/vet-medical-history-api/pkg/helpers"
	"log"
	"net/http"
)

func TokenValid(r *http.Request) (string, error) {
	tokenString := r.Header.Get("Authorization")

	client, err := helpers.FirebaseApp.Auth(context.Background())
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalf("error getting Auth client: %v\n", err)
		return "", err
	}

	token, err := client.VerifyIDToken(context.Background(), tokenString)
	if err != nil {
		return "", err
	}

	return token.UID, nil
}
