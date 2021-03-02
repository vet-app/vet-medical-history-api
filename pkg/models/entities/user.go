package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/vet-app/vet-medical-history-api/pkg/models"
	"time"
)

type User struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"size:60;not null" validate:"required"`
	Address   string    `json:"address" gorm:"size:100;not null"`
	Phone     string    `json:"phone_number" gorm:"column:phone_number;size:30;not null"`
	Email     string    `json:"email" gorm:"size:100;not null" validate:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func GetAllUsers() (*[]User, error) {
	var users []User
	err := models.DB.Debug().Model(&User{}).Limit(100).Find(&users).Error

	if err != nil {
		return &[]User{}, err
	}

	return &users, nil
}

func GetUserByID(id string) (*User, error) {
	var user User
	err := models.DB.Debug().Model(&User{}).Where("id = ?", id).Take(&user).Error

	if err != nil {
		return &User{}, err
	}

	return &user, nil
}

func CreateUser(user User) (string, error) {
	v := validator.New()
	err := v.Struct(user)
	if err != nil {
		return "", err
	}

	err = models.DB.Debug().Model(&User{}).Create(&user).Error
	if err != nil {
		return "", err
	}

	return user.ID, nil
}

func UpdateUser(id string, user User) error {
	err := models.DB.Debug().Model(&User{}).Where("id = ?", id).Updates(
		User{
			Name:      user.Name,
			Address:   user.Address,
			Phone:     user.Phone,
			Email:     user.Email,
			UpdatedAt: time.Now(),
		},
	).Error

	if err != nil {
		return err
	}

	return nil
}
