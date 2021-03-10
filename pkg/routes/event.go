package routes

import (
	"github.com/vet-app/vet-medical-history-api/pkg/controllers/events"
	"github.com/vet-app/vet-medical-history-api/pkg/middlewares"
)

var EventRoutes = Routes{
	Route{
		"Get Records By Pet",
		"GET",
		"/pet/{id}",
		middlewares.SetMiddlewareJSON(events.GetEventsByPet),
	},
	Route{
		"Get Record By ID",
		"GET",
		"/{id}",
		middlewares.SetMiddlewareJSON(events.GetEventsByID),
	},
	Route{
		"Create Record",
		"POST",
		"/",
		middlewares.SetMiddlewareJSON(events.CreateEvents),
	},
	Route{
		"Update Record",
		"PUT",
		"/{id}",
		middlewares.SetMiddlewareJSON(events.UpdateEvents),
	},
	Route{
		"Get Record By Pet And Tag",
		"GET",
		"/tags/pet/{id}",
		middlewares.SetMiddlewareJSON(events.GetEventsByPetAndType),
	},
}
