package events

import (
	"github.com/segmentio/ksuid"
	"github.com/vet-app/vet-medical-history-api/pkg/models"
	"github.com/vet-app/vet-medical-history-api/pkg/models/pets"
	"time"
)

type Event struct {
	ID          string    `json:"id" gorm:"primary_key"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	NextDate    time.Time `json:"next_date"`
	PetID       uint64    `json:"pet_id,omitempty" gorm:"REFERENCES pets(id)"`
	Pet         pets.Pet  `json:"pet,omitempty" gorm:"association_autoupdate:false;association_autocreate:false"`
	VetstoreID  string    `json:"vetstore_id,omitempty"`
	RecordID    string    `json:"record_id,omitempty" gorm:"REFERENCES records(id)"`
	Record      Record    `json:"record,omitempty" gorm:"association_autoupdate:false;association_autocreate:false"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func GetEventsByPet(petId string) (*[]Event, error) {
	var records []Event
	var result []Event
	err := models.DB.Debug().Model(&Event{}).
		Joins("JOIN \"pets\" p ON p.id = events.pet_id").
		Where("p.id = ?", petId).
		Find(&records).Error

	if err != nil {
		return &[]Event{}, err
	}

	if len(records) > 0 {
		for i, _ := range records {
			record, err := addEventData(records[i])
			result = append(result, *record)

			if err != nil {
				return &[]Event{}, err
			}
		}
	}

	return &result, nil
}

func GetEventsByID(id string) (*Event, error) {
	var event Event
	err := models.DB.Debug().Model(&Event{}).Where("id = ?", id).Take(&event).Error

	if err != nil {
		return &Event{}, err
	}

	if event.ID != "" {
		eventData, err := addEventData(event)

		if err != nil {
			return &Event{}, err
		}

		return eventData, nil
	}

	return &event, nil
}

func CreateEvents(event Event) error {
	event.ID = ksuid.New().String()
	err := models.DB.Debug().Model(&Event{}).Create(&event).Error

	if err != nil {
		return err
	}

	return nil
}

func UpdateEvents(id string, event Event) error {
	err := models.DB.Debug().Model(&Event{}).Where("id = ?", id).Updates(
		map[string]interface{}{
			"Title":       event.Title,
			"Description": event.Description,
			"StartDate":   event.StartDate,
			"NextDate":    event.NextDate,
			"PetID":       event.PetID,
			"VetstoreID":  event.VetstoreID,
			"RecordID":    event.RecordID,
			"UpdatedAt":   time.Now(),
		},
	).Error

	if err != nil {
		return err
	}

	return nil
}

func GetEventsByPetAndType(petId string, tag string) (*[]Event, error) {
	var events []Event
	var result []Event

	recordType, err := GetRecordTypeByTag(tag)

	if err != nil {
		return &[]Event{}, err
	}

	if recordType != nil {
		err := models.DB.Debug().Model(&Event{}).
			Joins("JOIN \"pets\" p ON p.id = events.pet_id").
			Joins("JOIN \"records\" r ON r.id = events.record_id").
			Joins("JOIN \"record_types\" t ON t.id = r.record_type_id").
			Where("p.id = ? AND t.id = ?", petId, recordType.ID).
			Find(&events).Error

		if err != nil {
			return &[]Event{}, err
		}
	}

	if len(events) > 0 {
		for i, _ := range events {
			event, err := addEventData(events[i])
			result = append(result, *event)

			if err != nil {
				return &[]Event{}, err
			}
		}
	}

	return &result, nil
}

func addEventData(event Event) (*Event, error) {
	err := models.DB.Debug().Model(&Event{}).Where("id = ?", event.PetID).
		Take(&event.Pet).Error

	if err != nil {
		return &Event{}, err
	}

	err = models.DB.Debug().Model(&Event{}).Where("id = ?", event.RecordID).
		Take(&event.Record).Error

	if err != nil {
		return &Event{}, err
	}

	err = models.DB.Debug().Model(&Event{}).Where("id = ?", event.Pet.UserID).
		Take(&event.Pet.User).Error

	if err != nil {
		return &Event{}, err
	}

	err = models.DB.Debug().Model(&Event{}).Where("id = ?", event.Pet.BreedID).
		Take(&event.Pet.Breed).Error

	if err != nil {
		return &Event{}, err
	}

	err = models.DB.Debug().Model(&Event{}).Where("id = ?", event.Pet.Breed.SpecieID).
		Take(&event.Pet.Breed.Specie).Error

	if err != nil {
		return &Event{}, err
	}

	event.PetID = 0
	event.RecordID = ""
	event.Pet.UserID = ""
	event.Pet.BreedID = ""
	event.Pet.Breed.SpecieID = ""
	return &event, nil
}
