package routes

import (
	"github.com/vet-app/vet-medical-history-api/pkg/controllers"
	"github.com/vet-app/vet-medical-history-api/pkg/middlewares"
)

var UsersRoutes = Routes{
	Route{
		"Get Users",
		"GET",
		"/",
		middlewares.SetMiddlewareJSON(controllers.GetUsers),
	},
	Route{
		"Get User By ID",
		"GET",
		"/{id}",
		middlewares.SetMiddlewareJSON(controllers.GetUserByID),
	},
	Route{
		"Create User",
		"POST",
		"/",
		middlewares.SetMiddlewareJSON(controllers.CreateUser),
	},
	Route{
		"Update User",
		"PUT",
		"/",
		middlewares.SetMiddlewareJSON(controllers.UpdateUser),
	},
}
