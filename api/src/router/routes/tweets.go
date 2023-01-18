package routes

import (
	"api/src/controllers"
	"net/http"
)

var routeTweets = []Route{
	{
		Uri:      "/tweets",
		Method:   http.MethodPost,
		Function: controllers.CreateTweet,
		NeedAuth: true,
	},
	{
		Uri:      "/tweets",
		Method:   http.MethodGet,
		Function: controllers.FindTweets,
		NeedAuth: true,
	},
	{
		Uri:      "/tweets/{tweetId}",
		Method:   http.MethodGet,
		Function: controllers.FindTweet,
		NeedAuth: true,
	},
	{
		Uri:      "/tweets/{tweetId}",
		Method:   http.MethodPut,
		Function: controllers.UpdateTweet,
		NeedAuth: true,
	},
	{
		Uri:      "/tweets/{tweetId}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteTweet,
		NeedAuth: true,
	},
}
