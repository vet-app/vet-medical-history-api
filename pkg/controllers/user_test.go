package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"github.com/vet-app/vet-medical-history-api/pkg/mocks"
	"github.com/vet-app/vet-medical-history-api/pkg/models/entities"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var muxVars = map[string]string{
	"id": "2",
}

var testCases = []struct {
	name    string
	isError bool
	status  int
}{
	{name: "Success procedure", isError: false, status: http.StatusOK},
	{name: "Server error", isError: true, status: http.StatusInternalServerError},
}

func TestGetUsers(t *testing.T) {
	mocks.SetupRepository()

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Errorf("On HTTP new request: %v\n", err)
	}
	rr := httptest.NewRecorder()
	GetUsers(rr, req)

	var users []entities.User
	err = json.Unmarshal([]byte(rr.Body.String()), &users)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "[]entities.User", reflect.TypeOf(users).String())
}

func TestGetUserByID(t *testing.T) {
	mocks.SetupRepository()

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.isError {
				mocket.Catcher.Attach([]*mocket.FakeResponse{
					{
						Pattern:  "SELECT * FROM \"users\"  WHERE",
						Response: mocks.MockReply,
						Once:     false,
					},
				})
			}

			req, err := http.NewRequest("GET", "users/3", nil)
			if err != nil {
				t.Errorf("On HTTP new request: %v\n", err)
			}
			req = mux.SetURLVars(req, muxVars)
			rr := httptest.NewRecorder()
			GetUserByID(rr, req)

			var user entities.User
			err = json.Unmarshal([]byte(rr.Body.String()), &user)
			if err != nil {
				log.Fatalf("Cannot convert to json: %v\n", err)
			}

			assert.Equal(t, tt.status, rr.Code)
			mocket.Catcher.Reset()
		})
	}
}

func TestCreateUser(t *testing.T) {
	mocks.SetupRepository()

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var jsonStr []byte
			jsonStr, _ = json.Marshal(mocks.MockReply[0])
			if tt.isError {
				jsonStr, _ = json.Marshal(mocks.MockError)
			}

			req, err := http.NewRequest("POST", "users/", bytes.NewBuffer(jsonStr))
			if err != nil {
				t.Errorf("On HTTP new request: %v\n", err)
			}
			req.URL.Query().Set("uid", "1")
			rr := httptest.NewRecorder()
			CreateUser(rr, req)

			assert.Equal(t, tt.status, rr.Code)
			mocket.Catcher.Reset()
		})
	}
}
