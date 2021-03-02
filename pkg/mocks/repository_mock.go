package mocks

import (
	"github.com/jinzhu/gorm"
	mocket "github.com/selvatico/go-mocket"
	"github.com/vet-app/vet-medical-history-api/pkg/models"
	"log"
	"time"
)

var MockReply = []map[string]interface{}{
	{
		"id":           "1",
		"name":         "John Smith",
		"address":      "Cra 163 No. 34 - 67",
		"phone_number": "408-237-2345",
		"email":        "john.smith@example.com",
		"created_at":   time.Time{},
		"updated_at":   time.Time{},
	},
	{
		"id":           "2",
		"name":         "John Smith",
		"address":      "Cra 163 No. 34 - 67",
		"phone_number": "408-237-2345",
		"email":        "john.smith@example.com",
		"created_at":   time.Time{},
		"updated_at":   time.Time{},
	},
	{
		"id":           "3",
		"name":         "John Smith",
		"address":      "Cra 163 No. 34 - 67",
		"phone_number": "408-237-2345",
		"email":        "john.smith@example.com",
		"created_at":   time.Time{},
		"updated_at":   time.Time{},
	},
}

var MockError = map[string]interface{}{
	"name":         123,
	"phone_number": "408-237-2345",
	"created_at":   "",
	"updated_at":    time.Time{},
}

func SetupRepository() {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	db, err := gorm.Open(mocket.DriverName, "mock_host")
	if err != nil {
		log.Println("Cannot connect to Postgres database")
		log.Fatal("Connection error:", err)
	}

	models.DB = db
}
