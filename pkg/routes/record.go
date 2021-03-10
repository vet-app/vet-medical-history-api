package routes

import (
	"github.com/vet-app/vet-medical-history-api/pkg/controllers/events"
	"github.com/vet-app/vet-medical-history-api/pkg/middlewares"
)

var RecordTypesRoutes = Routes{
	Route{
		"Search Record type By Tag",
		"GET",
		"/tag/",
		middlewares.SetMiddlewareJSON(events.SearchRecordTypeByTag),
	},
	Route{
		"Get Record Types",
		"GET",
		"/",
		middlewares.SetMiddlewareJSON(events.GetAllRecordTypes),
	},
	Route{
		"Get Record type By ID",
		"GET",
		"/{id}",
		middlewares.SetMiddlewareJSON(events.GetRecordTypeByID),
	},
	Route{
		"Create Record type",
		"POST",
		"/",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(events.CreateRecordType)),
	},
	Route{
		"Update Record type",
		"PUT",
		"/{id}",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(events.UpdateRecordType)),
	},
	Route{
		"Delete Record type",
		"DELETE",
		"/{id}",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(events.DeleteRecordType)),
	},
}

var RecordRoutes = Routes{
	Route{
		"Get All Records",
		"GET",
		"/",
		middlewares.SetMiddlewareJSON(events.GetAllRecords),
	},
	Route{
		"Get Record By ID",
		"GET",
		"/{id}",
		middlewares.SetMiddlewareJSON(events.GetRecordByID),
	},
	Route{
		"Create Record",
		"POST",
		"/",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(events.CreateRecord)),
	},
	Route{
		"Update Record",
		"PUT",
		"/{id}",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(events.UpdateRecord)),
	},
	Route{
		"Delete Record",
		"DELETE",
		"/{id}",
		middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(events.DeleteRecord)),
	},
}
