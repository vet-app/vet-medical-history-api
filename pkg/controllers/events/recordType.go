package events

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vet-app/vet-medical-history-api/pkg/models/events"
	"github.com/vet-app/vet-medical-history-api/pkg/responses"
	"net/http"
)

func GetAllRecordTypes(w http.ResponseWriter, r *http.Request) {
	types, err := events.GetAllRecordTypes()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, types)
}

func GetRecordTypeByID(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	recordType, err := events.GetRecordTypeByID(params["id"])
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, recordType)
}

func CreateRecordType(w http.ResponseWriter, r *http.Request) {
	var recordType events.RecordType
	_ = json.NewDecoder(r.Body).Decode(&recordType)
	err := events.CreateRecordType(recordType)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response := responses.RequestResponse{
		Response: "Categoria creada satisfactoriamente",
	}
	responses.JSON(w, http.StatusOK, response)
}

func UpdateRecordType(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	var recordType events.RecordType
	id := params["id"]
	_ = json.NewDecoder(r.Body).Decode(&recordType)
	err := events.UpdateRecordType(id, recordType)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response := responses.RequestResponse{
		Response: "Categoria editada satisfactoriamente",
	}
	responses.JSON(w, http.StatusOK, response)
}

func DeleteRecordType(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	id := params["id"]
	err := events.DeleteRecordType(id)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response := responses.RequestResponse{
		Response: "Categoria eliminada satisfactoriamente",
	}
	responses.JSON(w, http.StatusOK, response)
}

func SearchRecordTypeByTag(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("tag")
	recordType, err := events.GetRecordTypeByTag(id)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, recordType)
}
