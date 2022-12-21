package routes

import "net/http"

//Route represents all routes in the aplication
type Route struct {
	Uri      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	NeedAuth bool
}
