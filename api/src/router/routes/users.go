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
		NeedAuth: true,
	},
	{
		Uri:      "/users/{userId}",
		Method:   http.MethodGet,
		Function: controllers.GetUser,
		NeedAuth: true,
	},
	{
		Uri:      "/users/{userId}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		NeedAuth: true,
	},
	{
		Uri:      "/users/{userId}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
		NeedAuth: true,
	},
	{
		Uri:      "/users/{userId}/follow",
		Method:   http.MethodPost,
		Function: controllers.FollowUser,
		NeedAuth: true,
	},
}
