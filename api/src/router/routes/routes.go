package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all routes in the aplication
type Route struct {
	Uri      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	NeedAuth bool
}

// Configure configures all routes inside router
func Configure(r *mux.Router) *mux.Router {
	routes := routeUsers
	routes = append(routes, LoginRoute)

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}
	return r
}
