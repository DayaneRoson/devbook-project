package routes

import (
	"api/src/controllers"
	"net/http"
)

var LoginRoute = Route{
	Uri:      "/login",
	Method:   http.MethodPost,
	Function: controllers.Login,
	NeedAuth: false,
}
