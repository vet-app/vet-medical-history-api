package events

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vet-app/vet-medical-history-api/pkg/models/events"
	"github.com/vet-app/vet-medical-history-api/pkg/responses"
	"net/http"
)

func GetEventsByPet(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	petRecords, err := events.GetEventsByPet(params["id"])
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, petRecords)
}

func GetEventsByID(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	record, err := events.GetEventsByID(params["id"])
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, record)
}

func CreateEvents(w http.ResponseWriter, r *http.Request) {
	var record events.Event
	_ = json.NewDecoder(r.Body).Decode(&record)
	err := events.CreateEvents(record)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response := responses.RequestResponse{
		Response: "Evento creado satisfactoriamente",
	}
	responses.JSON(w, http.StatusOK, response)
}

func UpdateEvents(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	var record events.Event
	id := params["id"]
	_ = json.NewDecoder(r.Body).Decode(&record)
	err := events.UpdateEvents(id, record)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response := responses.RequestResponse{
		Response: "Evento editado satisfactoriamente",
	}
	responses.JSON(w, http.StatusOK, response)
}

func GetEventsByPetAndType(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	tag := r.URL.Query().Get("type")
	petRecords, err := events.GetEventsByPetAndType(params["id"], tag)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if len(*petRecords) == 0 {
		responses.JSON(w, http.StatusOK, []map[string]interface{}{})
		return
	}

	responses.JSON(w, http.StatusOK, petRecords)
}
