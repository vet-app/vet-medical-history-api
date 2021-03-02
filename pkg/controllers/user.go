package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vet-app/vet-medical-history-api/pkg/models/entities"
	"github.com/vet-app/vet-medical-history-api/pkg/responses"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := entities.GetAllUsers()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	id := params["id"]
	user, err := entities.GetUserByID(id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	user.ID = r.URL.Query().Get("uid")
	_ = json.NewDecoder(r.Body).Decode(&user)
	id, err := entities.CreateUser(user)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response := responses.RequestResponse{
		Response: "Usuario creado satisfactoriamente",
		ID:       id,
	}
	responses.JSON(w, http.StatusOK, response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	id := r.URL.Query().Get("uid")
	_ = json.NewDecoder(r.Body).Decode(&user)
	err := entities.UpdateUser(id, user)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response := responses.RequestResponse{
		Response: "Usuario editado satisfactoriamente",
	}
	responses.JSON(w, http.StatusOK, response)
}
