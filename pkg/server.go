package pkg

import (
	"github.com/getsentry/sentry-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/vet-app/vet-medical-history-api/pkg/helpers"
	"github.com/vet-app/vet-medical-history-api/pkg/models"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type Server struct {
	Router *mux.Router
}

func (server *Server) Initialize(dbHost, dbPort, dbUser, dbPassword, dbName, ssl string) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
		Environment: os.Getenv("ENV"),
		Release: "vet-history@1.0.0",
		Debug: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go models.DBConnection(wg, dbHost, dbPort, dbUser, dbPassword, dbName, ssl)
	wg.Wait()

	helpers.FirebaseConnection()

	server.Router = mux.NewRouter().StrictSlash(true)
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	log.Println("Server started successfully")
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"origin", "accept", "X-Requested-With", "Content-Type", "Authorization", "authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS", "HEAD"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
	)
	log.Fatal(http.ListenAndServe(addr, cors(server.Router)))
}
