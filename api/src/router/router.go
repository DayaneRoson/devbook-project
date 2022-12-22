package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate generates a router and returns it
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
