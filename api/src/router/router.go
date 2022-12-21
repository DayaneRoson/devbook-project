package router

import "github.com/gorilla/mux"

//Generate generates a router and returns it
func Generate() *mux.Router {
	return mux.NewRouter()
}
