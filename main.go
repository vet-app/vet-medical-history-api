package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("$PORT not set")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", HealthCheck)
	http.Handle("/", r)
	fmt.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

func HealthCheck(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintln(w, "Alive!")
}
