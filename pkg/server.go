package pkg

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/vet-app/vet-medical-history-api/pkg/helpers"
	"github.com/vet-app/vet-medical-history-api/pkg/models"
	"log"
	"net/http"
	"sync"
)

type Server struct {
	Router *mux.Router
}

func (server *Server) Initialize(dbHost, dbPort, dbUser, dbPassword, dbName, ssl string) {
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
