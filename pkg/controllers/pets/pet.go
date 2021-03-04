package pets

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vet-app/vet-medical-history-api/pkg/helpers"
	"github.com/vet-app/vet-medical-history-api/pkg/models/pets"
	"github.com/vet-app/vet-medical-history-api/pkg/responses"
	"net/http"
	"strconv"
)

func GetPetsByUser(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	p, err := pets.GetPetsByUser(params["id"])
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, p)
}

func GetPetByID(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 0, 64)
	pet, err := pets.GetPetByID(id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, pet)
}

func CreatePet(w http.ResponseWriter, r *http.Request) {
	var pet pets.Pet
	_ = json.NewDecoder(r.Body).Decode(&pet)

	pet.UserID = r.URL.Query().Get("uid")
	pet.BreedID = pet.Breed.ID

	filename, id, err := pets.CreatePet(pet)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response := responses.RequestResponse{
		Response: "Datos guardados satisfactoriamente",
		ID: strconv.FormatUint(id, 10),
		Filename: filename,
	}
	responses.JSON(w, http.StatusOK, response)
}

func UpdatePet(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	var pet pets.Pet
	id, _ := strconv.ParseUint(params["id"], 0, 64)
	_ = json.NewDecoder(r.Body).Decode(&pet)
	err := pets.UpdatePet(id, pet)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response := responses.RequestResponse{
		Response: "Mascota editada satisfactoriamente",
	}
	responses.JSON(w, http.StatusOK, response)
}

func UploadPetPhoto(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	file, _, err := r.FormFile("profile_picture")
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer file.Close()

	filename := r.FormValue("filename")

	err = helpers.StoreFile(file, "pets/", filename)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response := responses.RequestResponse{
		Response: "Imagen de mascota guardada satisfactoriamente",
	}
	responses.JSON(w, http.StatusOK, response)
}

func UpdatePetPhoto(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 0, 64)

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	file, _, err := r.FormFile("profile_picture")
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer file.Close()

	filemime := r.FormValue("mime_type")

	filename, err := pets.UpdatePetPhoto(id, filemime)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if filemime != "" {
		err = helpers.StoreFile(file, "pets/", filename)

		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
	}

	response := responses.RequestResponse{
		Response: "Imagen de mascota editada satisfactoriamente",
	}
	responses.JSON(w, http.StatusOK, response)
}

func SearchPetID(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	_ = json.NewDecoder(r.Body).Decode(&data)
	p, err := pets.SearchPetID(data["keyword"].(uint64))

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, p)
}
