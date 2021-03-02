package database

import (
	"github.com/vet-app/vet-medical-history-api/pkg/models"
	"github.com/vet-app/vet-medical-history-api/pkg/models/entities"
	"log"
)

func MigrateDatabase() {
	var dbModels = []interface{}{&entities.User{}}

	err := models.DB.AutoMigrate(dbModels...)
	if err != nil {
		log.Fatal("Error on models migration:", err)
	}
}
