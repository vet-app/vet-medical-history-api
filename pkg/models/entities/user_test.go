package entities

import (
	"encoding/json"
	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"github.com/vet-app/vet-medical-history-api/pkg/mocks"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestGetAllUsers(t *testing.T) {
	mocks.SetupRepository()

	mocket.Catcher.Attach([]*mocket.FakeResponse{
		{
			Pattern:  "SELECT * FROM \"users\"",
			Response: mocks.MockReply,
			Once:     false,
		},
	})

	users, err := GetAllUsers()
	if err != nil {
		t.Errorf("Error getting users: %v\n", err)
		return
	}

	if assert.NotNil(t, users) {
		assert.Equal(t, 3, len(*users))
		assert.Equal(t,
			User{
				ID:        "2",
				Name:      "John Smith",
				Address:   "Cra 163 No. 34 - 67",
				Phone:     "408-237-2345",
				Email:     "john.smith@example.com",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			}, (*users)[1])
	}
}

func TestGetUserByID(t *testing.T) {
	mocks.SetupRepository()

	mocket.Catcher.Attach([]*mocket.FakeResponse{
		{
			Pattern:  "SELECT * FROM \"users\" WHERE",
			Response: mocks.MockReply,
			Once:     false,
		},
	})

	user, err := GetUserByID("3")
	if err != nil {
		t.Errorf("Error getting users: %v\n", err)
		return
	}

	if assert.NotNil(t, user) {
		assert.Equal(t, "3", user.ID)
	}
}

func TestCreateUser(t *testing.T) {
	mocks.SetupRepository()

	var mockedUser User
	jsonStr, _ := json.Marshal(mocks.MockReply[0])
	err := json.Unmarshal(jsonStr, &mockedUser)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}

	mockedId, _ := strconv.ParseInt(mockedUser.ID, 10, 64)
	mocket.Catcher.Reset().NewMock().WithQuery("INSERT INTO \"users\"").WithID(mockedId)
	returnedId, _ := CreateUser(mockedUser)

	assert.Equal(t, "1", returnedId)
}

func TestUpdateUser(t *testing.T) {
	mocks.SetupRepository()

	var mockedUser User
	jsonStr, _ := json.Marshal(mocks.MockReply[0])
	err := json.Unmarshal(jsonStr, &mockedUser)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}

	mockedId, _ := strconv.ParseInt(mockedUser.ID, 10, 64)
	mocket.Catcher.Reset().NewMock().WithQuery("UPDATE \"users\"").WithID(mockedId)
	err = UpdateUser("1", mockedUser)

	assert.Equal(t, nil, err)
}
