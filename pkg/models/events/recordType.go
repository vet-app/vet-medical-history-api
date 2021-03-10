package events

import (
	"github.com/segmentio/ksuid"
	"github.com/vet-app/vet-medical-history-api/pkg/models"
)

type RecordType struct {
	ID      string `json:"id,omitempty" gorm:"primary_key"`
	Name    string `json:"name,omitempty" gorm:"size:60;not null"`
	Tag     string `json:"tag,omitempty" gorm:"size:60"`
	Deleted bool   `json:"-"`
}

func GetAllRecordTypes() (*[]RecordType, error) {
	var recordTypes []RecordType
	err := models.DB.Debug().Model(&RecordType{}).
		Where("deleted = ?", false).
		Limit(100).
		Find(&recordTypes).Error

	if err != nil {
		return &[]RecordType{}, err
	}

	return &recordTypes, nil
}

func GetRecordTypeByID(id string) (*RecordType, error) {
	var recordType RecordType
	err := models.DB.Debug().Model(&RecordType{}).Where("id = ?", id).Take(&recordType).Error

	if err != nil {
		return &RecordType{}, err
	}

	return &recordType, nil
}

func CreateRecordType(recordType RecordType) error {
	recordType.ID = ksuid.New().String()
	recordType.Deleted = false
	err := models.DB.Debug().Model(&RecordType{}).Create(&recordType).Error

	if err != nil {
		return err
	}

	return nil
}

func UpdateRecordType(id string, recordType RecordType) error {
	err := models.DB.Debug().Model(&RecordType{}).Where("id = ?", id).Updates(
		map[string]interface{}{
			"Name":    recordType.Name,
			"Deleted": false,
		},
	).Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteRecordType(id string) error {
	err := models.DB.Debug().Model(&RecordType{}).Where("id = ?", id).Updates(
		RecordType{
			Deleted: true,
		},
	).Error

	if err != nil {
		return err
	}

	return nil
}

func GetRecordTypeByTag(tag string) (*RecordType, error) {
	var recordType RecordType
	err := models.DB.Debug().Model(&RecordType{}).Where("tag = ?", tag).Take(&recordType).Error

	if err != nil {
		return &RecordType{}, err
	}

	return &recordType, nil
}
