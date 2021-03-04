package middlewares

import (
	"errors"
	"github.com/vet-app/vet-medical-history-api/pkg/auth"
	"github.com/vet-app/vet-medical-history-api/pkg/responses"
	"net/http"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		uid, err := auth.TokenValid(r)

		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		q.Add("uid", uid)
		r.URL.RawQuery = q.Encode()
		next(w, r)
	}
}
