package models

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	"log"
	"sync"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func DBConnection(wg *sync.WaitGroup, dbHost, dbPort, dbUser, dbPassword, dbName, ssl string) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbName, dbPassword, ssl)

	defer wg.Done()

	database, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Println("Cannot connect to Postgres database")
		sentry.CaptureException(err)
	}

	log.Println("Connected to Postgres database")
	DB = database
}
