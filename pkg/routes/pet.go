package routes

import (
	"github.com/vet-app/vet-medical-history-api/pkg/controllers/pets"
	"github.com/vet-app/vet-medical-history-api/pkg/middlewares"
)

var SpeciesRoutes = Routes{
	Route{
		"Get Species",
		"GET",
		"/",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(pets.GetSpecies)),
	},
	Route{
		"Get Specie By ID",
		"GET",
		"/{id}",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(pets.GetSpecieByID)),
	},
	Route{
		"Create Specie",
		"POST",
		"/",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(pets.CreateSpecie)),
	},
	Route{
		"Update Specie",
		"PUT",
		"/{id}",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(pets.UpdateSpecie)),
	},
	Route{
		"Delete Species",
		"DELETE",
		"/{id}",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(pets.DeleteSpecie)),
	},
}

var BreedsRoutes = Routes{
	Route{
		"Get Breeds",
		"GET",
		"/",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(pets.GetBreeds)),
	},
	Route{
		"Get Breed By ID",
		"GET",
		"/{id}",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(pets.GetBreedByID)),
	},
	Route{
		"Create Breed",
		"POST",
		"/",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(pets.CreateBreed)),
	},
	Route{
		"Update Breed",
		"PUT",
		"/{id}",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(pets.UpdateBreed)),
	},
	Route{
		"Delete Breed",
		"DELETE",
		"/{id}",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(pets.DeleteBreed)),
	},
	Route{
		"Get Breed by Specie ID",
		"GET",
		"/specie/{id}",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(pets.GetBreedsBySpecie)),
	},
}

var PetsRoutes = Routes{
	Route{
		"Get Pets By User",
		"GET",
		"/user/{id}",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(pets.GetPetsByUser)),
	},
	Route{
		"Get Pet By ID",
		"GET",
		"/{id}",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(pets.GetPetByID)),
	},
	Route{
		"Create Pet",
		"POST",
		"/",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(pets.CreatePet)),
	},
	Route{
		"Update Pet",
		"PUT",
		"/{id}",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(pets.UpdatePet)),
	},
	Route{
		"Upload Pet Photo",
		"POST",
		"/uploadPhoto/",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(pets.UploadPetPhoto)),
	},
	Route{
		"Update Pet Photo",
		"PUT",
		"/updatePhoto/{id}",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(pets.UpdatePetPhoto)),
	},
	Route{
		"Search Products",
		"POST",
		"/search/",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(pets.SearchPetID)),
	},
}
