package routes

import (
	middlewares "api/src/middleware"
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
	routes = append(routes, routeTweets...)

	for _, route := range routes {
		if route.NeedAuth {
			r.HandleFunc(route.Uri, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}
	return r
}
