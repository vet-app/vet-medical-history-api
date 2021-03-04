package pkg

import (
	"github.com/gorilla/mux"
	"github.com/vet-app/vet-medical-history-api/pkg/routes"
)

func (server *Server) initializeRoutes() {
	m := make(map[*mux.Router]routes.Routes)
	apiDeclaration := "/api/v1"

	apiPrefix := server.Router.PathPrefix(apiDeclaration).Subrouter()

	userRouter := apiPrefix.PathPrefix("/users").Subrouter().StrictSlash(true)
	m[userRouter] = routes.UsersRoutes

	petRouter := apiPrefix.PathPrefix("/pets").Subrouter().StrictSlash(true)
	m[petRouter] = routes.PetsRoutes

	specieRouter := apiPrefix.PathPrefix("/species").Subrouter().StrictSlash(true)
	m[specieRouter] = routes.SpeciesRoutes

	breedRouter := apiPrefix.PathPrefix("/breeds").Subrouter().StrictSlash(true)
	m[breedRouter] = routes.BreedsRoutes

	for router, r := range m {
		setRoutes(router, r)
	}
}

func setRoutes(router *mux.Router, routes routes.Routes) {
	for _, route := range routes {
		handler := route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
}
