package pets

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vet-app/vet-medical-history-api/pkg/models/pets"
	"github.com/vet-app/vet-medical-history-api/pkg/responses"
	"net/http"
)

func GetSpecies(w http.ResponseWriter, r *http.Request) {
	species, err := pets.GetAllSpecies()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, species)
}

func GetSpecieByID(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	specie, err := pets.GetSpecieByID(params["id"])
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, specie)
}

func CreateSpecie(w http.ResponseWriter, r *http.Request) {
	var specie pets.Specie
	_ = json.NewDecoder(r.Body).Decode(&specie)
	err := pets.CreateSpecie(specie)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response := responses.RequestResponse{
		Response: "Especie animal creada satisfactoriamente",
	}
	responses.JSON(w, http.StatusOK, response)
}

func UpdateSpecie(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	var specie pets.Specie
	id := params["id"]
	_ = json.NewDecoder(r.Body).Decode(&specie)
	err := pets.UpdateSpecie(id, specie)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response := responses.RequestResponse{
		Response: "Especie animal editada satisfactoriamente",
	}
	responses.JSON(w, http.StatusOK, response)
}

func DeleteSpecie(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	id := params["id"]
	err := pets.DeleteSpecie(id)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response := responses.RequestResponse{
		Response: "Especie animal eliminada satisfactoriamente",
	}
	responses.JSON(w, http.StatusOK, response)
}
