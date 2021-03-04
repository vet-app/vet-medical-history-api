package pets

import (
	"github.com/vet-app/vet-medical-history-api/pkg/helpers"
	"github.com/vet-app/vet-medical-history-api/pkg/models"
	"github.com/vet-app/vet-medical-history-api/pkg/models/entities"
	"time"
)

type Pet struct {
	ID         string        `json:"id" gorm:"primary_key;autoIncrement;"`
	Name       string        `json:"name" gorm:"size:60;not null"`
	BornDate   time.Time     `json:"born_date" gorm:"not null"`
	Weight     string        `json:"weight" gorm:"size:10"`
	UserID     string        `json:"user_id,omitempty" gorm:"REFERENCES user(id)"`
	User       entities.User `json:"user" gorm:"association_autoupdate:false;association_autocreate:false"`
	BreedID    string        `json:"breed_id,omitempty" gorm:"REFERENCES breed(id)"`
	Breed      Breed         `json:"breed" gorm:"association_autoupdate:false;association_autocreate:false"`
	PictureURL string        `json:"picture_url" gorm:"column:picture_url;not null"`
	FileMime   string        `json:"mime_type" gorm:"column:mime_type"`
	CreatedAt  time.Time     `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time     `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

var petBucketUrl = "https://firebasestorage.googleapis.com/v0/b/firuapp-62849.appspot.com/o/pets%2F"

func GetPetsByUser(userId string) (*[]Pet, error) {
	var pets []Pet
	var result []Pet
	err := models.DB.Debug().Model(&Pet{}).
		Joins("JOIN \"user\" u ON u.id = pet.user_id").
		Where("u.id = ?", userId).
		Find(&pets).Error

	if err != nil {
		return &[]Pet{}, err
	}

	if len(pets) > 0 {
		for i, _ := range pets {
			pet, err := AddPetData(pets[i])
			result = append(result, *pet)

			if err != nil {
				return &[]Pet{}, err
			}
		}
		return &result, nil
	}

	return &pets, nil
}

func GetPetByID(id string) (*Pet, error) {
	var pet Pet
	err := models.DB.Debug().Model(&Pet{}).Where("id = ?", id).Take(&pet).Error

	if err != nil {
		return &Pet{}, err
	}

	if pet.ID != "" {
		mascot, err := AddPetData(pet)

		if err != nil {
			return &Pet{}, err
		}

		return mascot, nil
	}

	return &pet, nil
}

func CreatePet(pet Pet) (string, error) {
	filename := helpers.AddFilename(pet.FileMime)

	pet.PictureURL = petBucketUrl + filename + "?alt=media"

	err := models.DB.Debug().Model(&Pet{}).Create(&pet).Error

	if err != nil {
		return "", err
	}

	return filename, nil
}

func UpdatePet(id string, pet Pet) error {
	err := models.DB.Debug().Model(&Pet{}).Where("id = ?", id).Updates(
		map[string]interface{}{
			"Name":      pet.Name,
			"BornDate":  pet.BornDate,
			"Weight":    pet.Weight,
			"BreedID":   pet.BreedID,
			"UpdatedAt": time.Now(),
		},
	).Error

	if err != nil {
		return err
	}

	return nil
}

func UpdatePetPhoto(id string, fileMime string) (string, error) {
	filename := helpers.AddFilename(fileMime)
	pictureUrl := petBucketUrl + filename + "?alt=media"

	err := models.DB.Debug().Model(&Pet{}).Where("id = ?", id).Updates(
		Pet{
			PictureURL: pictureUrl,
			FileMime:   fileMime,
		},
	).Error

	if err != nil {
		return "", err
	}

	return filename, nil
}

func SearchPetID(id string) (*[]Pet, error) {
	var pets []Pet
	var result []Pet
	err := models.DB.Debug().Model(&Pet{}).Where("id LIKE ?", "%"+id+"%").
		Limit(100).Find(&pets).Error

	if err != nil {
		return &[]Pet{}, err
	}

	if len(pets) > 0 {
		for i, _ := range pets {
			pet, err := AddPetData(pets[i])
			result = append(result, *pet)

			if err != nil {
				return &[]Pet{}, err
			}
		}
		return &result, nil
	}

	return &pets, nil
}

func AddPetData(pet Pet) (*Pet, error) {
	err := models.DB.Debug().Model(&Pet{}).Where("id = ?", pet.UserID).
		Take(&pet.User).Error

	if err != nil {
		return &Pet{}, err
	}

	err = models.DB.Debug().Model(&Pet{}).Where("id = ?", pet.BreedID).
		Take(&pet.Breed).Error

	if err != nil {
		return &Pet{}, err
	}

	err = models.DB.Debug().Model(&Pet{}).Where("id = ?", pet.Breed.SpecieID).
		Take(&pet.Breed.Specie).Error

	if err != nil {
		return &Pet{}, err
	}

	pet.BreedID = ""
	pet.Breed.SpecieID = ""
	pet.UserID = ""
	return &pet, nil
}
