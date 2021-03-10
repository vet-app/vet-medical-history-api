package events

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vet-app/vet-medical-history-api/pkg/models/events"
	"github.com/vet-app/vet-medical-history-api/pkg/responses"
	"net/http"
)

func GetAllRecords(w http.ResponseWriter, r *http.Request) {
	petRecords, err := events.GetAllRecords()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, petRecords)
}

func GetRecordByID(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	record, err := events.GetRecordByID(params["id"])
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, record)
}

func CreateRecord(w http.ResponseWriter, r *http.Request) {
	var record events.Record
	_ = json.NewDecoder(r.Body).Decode(&record)
	err := events.CreateRecord(record)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response := responses.RequestResponse{
		Response: "Registro creado satisfactoriamente",
	}
	responses.JSON(w, http.StatusOK, response)
}

func UpdateRecord(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	var record events.Record
	id := params["id"]
	_ = json.NewDecoder(r.Body).Decode(&record)
	err := events.UpdateRecord(id, record)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response := responses.RequestResponse{
		Response: "Registro editado satisfactoriamente",
	}
	responses.JSON(w, http.StatusOK, response)
}

func DeleteRecord(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	id := params["id"]
	err := events.DeleteRecord(id)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response := responses.RequestResponse{
		Response: "Registro eliminado satisfactoriamente",
	}
	responses.JSON(w, http.StatusOK, response)
}
