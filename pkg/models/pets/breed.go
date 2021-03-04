package pets

import (
	"github.com/segmentio/ksuid"
	"github.com/vet-app/vet-medical-history-api/pkg/models"
)

type Breed struct {
	ID       string `json:"id,omitempty" gorm:"primary_key"`
	Name     string `json:"name,omitempty" gorm:"size:60;not null"`
	SpecieID string `json:"specie_id,omitempty" gorm:"REFERENCES specie(id)"`
	Specie   Specie `json:"specie" gorm:"association_autoupdate:false;association_autocreate:false"`
	Deleted  bool   `json:"-"`
}

func GetAllBreeds() (*[]Breed, error) {
	var breeds []Breed
	var result []Breed
	err := models.DB.Debug().Model(&Breed{}).
		Where("deleted = ?", false).
		Limit(100).
		Find(&breeds).Error

	if err != nil {
		return &[]Breed{}, err
	}

	if len(breeds) > 0 {
		for i, _ := range breeds {
			pet, err := addBreedData(breeds[i])
			result = append(result, *pet)

			if err != nil {
				return &[]Breed{}, err
			}
		}
	}

	return &breeds, nil
}

func GetBreedByID(id string) (*Breed, error) {
	var breed Breed
	err := models.DB.Debug().Model(&Breed{}).Where("id = ?", id).Take(&breed).Error

	if err != nil {
		return &Breed{}, err
	}

	result, err := addBreedData(breed)

	if err != nil {
		return &Breed{}, err
	}

	return result, nil
}

func CreateBreed(breed Breed) error {
	breed.ID = ksuid.New().String()
	breed.Deleted = false
	err := models.DB.Debug().Model(&Breed{}).Create(&breed).Error

	if err != nil {
		return err
	}

	return nil
}

func UpdateBreed(id string, breed Breed) error {
	err := models.DB.Debug().Model(&Breed{}).Where("id = ?", id).Updates(
		map[string]interface{}{
			"Name":     breed.Name,
			"SpecieID": breed.SpecieID,
			"Deleted":  false,
		},
	).Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteBreed(id string) error {
	err := models.DB.Debug().Model(&Breed{}).Where("id = ?", id).Updates(
		Breed{
			Deleted: true,
		},
	).Error

	if err != nil {
		return err
	}

	return nil
}

func GetBreedBySpecie(id string) (*[]Breed, error) {
	var breeds []Breed
	var result []Breed
	err := models.DB.Debug().Model(&Breed{}).
		Joins("JOIN \"species\" s ON s.id = breed.specie_id").
		Where("s.id = ?", id).Take(&breeds).Error

	if err != nil {
		return &[]Breed{}, err
	}

	if len(breeds) > 0 {
		for i, _ := range breeds {
			breed, err := addBreedData(breeds[i])
			result = append(result, *breed)

			if err != nil {
				return &[]Breed{}, err
			}
		}
	}

	return &result, nil
}

func addBreedData(breed Breed) (*Breed, error) {
	err := models.DB.Debug().Model(&Breed{}).Where("id = ?", breed.SpecieID).
		Take(&breed.Specie).Error

	if err != nil {
		return &Breed{}, err
	}

	breed.SpecieID = ""
	return &breed, nil
}
