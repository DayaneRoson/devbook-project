package routes

import (
	"api/src/controllers"
	"net/http"
)

var routeUsers = []Route{
	{
		Uri:      "/users",
		Method:   http.MethodPost,
		Function: controllers.CreateUser,
		NeedAuth: false,
	},
	{
		Uri:      "/users",
		Method:   http.MethodGet,
		Function: controllers.GetUsers,
		NeedAuth: false,
	},
	{
		Uri:      "/users/{userId}",
		Method:   http.MethodGet,
		Function: controllers.GetUser,
		NeedAuth: false,
	},
	{
		Uri:      "/users/{userId}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		NeedAuth: false,
	},
	{
		Uri:      "/users/{userId}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
		NeedAuth: false,
	},
}
