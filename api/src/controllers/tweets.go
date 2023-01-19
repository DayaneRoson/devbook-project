package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateTweet creates a tweet
func CreateTweet(w http.ResponseWriter, r *http.Request) {
	userId, error := authentication.ExtractUserId(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	body, error := io.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	var tweet models.Tweet
	if error = json.Unmarshal(body, &tweet); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}
	tweet.AuthorId = userId
	if error = tweet.Prepare(); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connection()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewTweetRepository(db)
	tweet.ID, error = repository.Create(tweet)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusCreated, tweet)
}

// FindTweets brings tweets that would show on the user feed
func FindTweets(w http.ResponseWriter, r *http.Request) {

}

// FindTweet brings an specific user
func FindTweet(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	tweetId, error := strconv.ParseUint(parameters["tweetId"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}
	db, error := database.Connection()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()
	repository := repositories.NewTweetRepository(db)
	tweet, error := repository.FindById(tweetId)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusOK, tweet)
}

// UpdateTweet updates a tweet
func UpdateTweet(w http.ResponseWriter, r *http.Request) {

}

// DeleteTweet deletes a tweet
func DeleteTweet(w http.ResponseWriter, r *http.Request) {

}
