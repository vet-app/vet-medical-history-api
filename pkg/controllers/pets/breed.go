package pets

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vet-app/vet-medical-history-api/pkg/models/pets"
	"github.com/vet-app/vet-medical-history-api/pkg/responses"
	"net/http"
)

func GetBreeds(w http.ResponseWriter, r *http.Request) {
	breeds, err := pets.GetAllBreeds()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, breeds)
}

func GetBreedByID(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	breed, err := pets.GetBreedByID(params["id"])
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, breed)
}

func CreateBreed(w http.ResponseWriter, r *http.Request) {
	var breed pets.Breed
	_ = json.NewDecoder(r.Body).Decode(&breed)
	err := pets.CreateBreed(breed)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response := responses.RequestResponse{
		Response: "Raza creada satisfactoriamente",
	}
	responses.JSON(w, http.StatusOK, response)
}

func UpdateBreed(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	var breed pets.Breed
	id := params["id"]
	_ = json.NewDecoder(r.Body).Decode(&breed)
	err := pets.UpdateBreed(id, breed)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response := responses.RequestResponse{
		Response: "Raza editada satisfactoriamente",
	}
	responses.JSON(w, http.StatusOK, response)
}

func DeleteBreed(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	id := params["id"]
	err := pets.DeleteBreed(id)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response := responses.RequestResponse{
		Response: "Raza eliminada satisfactoriamente",
	}
	responses.JSON(w, http.StatusOK, response)
}

func GetBreedsBySpecie(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	breeds, err := pets.GetBreedBySpecie(params["id"])
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, breeds)
}
