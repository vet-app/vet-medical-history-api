package events

import (
	"github.com/segmentio/ksuid"
	"github.com/vet-app/vet-medical-history-api/pkg/models"
	"github.com/vet-app/vet-medical-history-api/pkg/models/pets"
)

type Record struct {
	ID           string      `json:"id,omitempty" gorm:"primary_key"`
	Title        string      `json:"title,omitempty" gorm:"size:60;not null"`
	SpecieID     string      `json:"specie_id,omitempty" gorm:"REFERENCES species(id)"`
	Specie       pets.Specie `json:"specie" gorm:"association_autoupdate:false;association_autocreate:false"`
	RecordTypeID string      `json:"record_type_id,omitempty" gorm:"REFERENCES record_types(id)"`
	RecordType   RecordType  `json:"record_type,omitempty" gorm:"association_autoupdate:false;association_autocreate:false"`
	Deleted      bool        `json:"-"`
}

func GetAllRecords() (*[]Record, error) {
	var records []Record
	err := models.DB.Debug().Model(&Record{}).
		Where("deleted = ?", false).
		Limit(100).
		Find(&records).Error

	if err != nil {
		return &[]Record{}, err
	}

	return &records, nil
}

func GetRecordByID(id string) (*Record, error) {
	var record Record
	err := models.DB.Debug().Model(&Record{}).Where("id = ?", id).Take(&record).Error

	if err != nil {
		return &Record{}, err
	}

	return &record, nil
}

func CreateRecord(record Record) error {
	record.ID = ksuid.New().String()
	record.Deleted = false
	err := models.DB.Debug().Model(&Record{}).Create(&record).Error

	if err != nil {
		return err
	}

	return nil
}

func UpdateRecord(id string, record Record) error {
	err := models.DB.Debug().Model(&Record{}).Where("id = ?", id).Updates(
		map[string]interface{}{
			"Title":    record.Title,
			"RecordTypeID":    record.RecordTypeID,
			"SpecieID": record.SpecieID,
			"Deleted": false,
		},
	).Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteRecord(id string) error {
	err := models.DB.Debug().Model(&Record{}).Where("id = ?", id).Updates(
		Record{
			Deleted: true,
		},
	).Error

	if err != nil {
		return err
	}

	return nil
}
