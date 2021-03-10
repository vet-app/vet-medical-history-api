package pets

import (
	"github.com/segmentio/ksuid"
	"github.com/vet-app/vet-medical-history-api/pkg/models"
)

type Specie struct {
	ID      string `json:"id,omitempty" gorm:"primary_key"`
	Name    string `json:"name,omitempty" gorm:"size:60;not null"`
	Deleted bool   `json:"-"`
}

func GetAllSpecies() (*[]Specie, error) {
	var species []Specie
	err := models.DB.Debug().Model(&Specie{}).
		Where("deleted = ?", false).
		Limit(100).
		Find(&species).Error

	if err != nil {
		return &[]Specie{}, err
	}

	return &species, nil
}

func GetSpecieByID(id string) (*Specie, error) {
	var specie Specie
	err := models.DB.Debug().Model(&Specie{}).Where("id = ?", id).Take(&specie).Error

	if err != nil {
		return &Specie{}, err
	}

	return &specie, nil
}

func CreateSpecie(specie Specie) error {
	specie.ID = ksuid.New().String()
	specie.Deleted = false
	err := models.DB.Debug().Model(&Specie{}).Create(&specie).Error

	if err != nil {
		return err
	}

	return nil
}

func UpdateSpecie(id string, specie Specie) error {
	err := models.DB.Debug().Model(&Specie{}).Where("id = ?", id).Updates(
		map[string]interface{}{
			"Title":    specie.Name,
			"Deleted": false,
		},
	).Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteSpecie(id string) error {
	err := models.DB.Debug().Model(&Specie{}).Where("id = ?", id).Updates(
		Specie{
			Deleted: true,
		},
	).Error

	if err != nil {
		return err
	}

	return nil
}
