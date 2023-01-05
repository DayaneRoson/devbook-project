package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser creates an user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, error := io.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}
	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.Error(w, http.StatusUnsupportedMediaType, error)
		return
	}
	if error = user.Prepare("register"); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connection()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user.ID, error = repository.Create(user)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusCreated, user)
}

// GetUsers fetches all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, error := database.Connection()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()
	repository := repositories.NewUserRepository(db)
	users, error := repository.Find(nameOrNick)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

// GetUser fetches an user
func GetUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userId, error := strconv.ParseUint(parameters["userId"], 10, 32)
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

	repository := repositories.NewUserRepository(db)
	user, error := repository.FindById(userId)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusOK, user)

}

// UpdateUser updates an user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, error := strconv.ParseUint(parameters["userId"], 10, 32)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	tokenUserId, error := authentication.ExtractUserId(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	if tokenUserId != userId {
		responses.Error(w, http.StatusForbidden, errors.New("forbidden access"))
		return
	}

	body, error := io.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}
	var user models.User
	if error = json.Unmarshal(body, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = user.Prepare("update"); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connection()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if error = repository.Update(userId, user); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// DeleteUser deletes an user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, error := strconv.ParseUint(parameters["userId"], 10, 32)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	tokenUserId, error := authentication.ExtractUserId(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	if tokenUserId != userId {
		responses.Error(w, http.StatusForbidden, errors.New("forbidden access"))
		return
	}

	db, error := database.Connection()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if error = repository.Delete(userId); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

// FollowUser allows a user to follow another
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, error := authentication.ExtractUserId(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	parameters := mux.Vars(r)
	followedUserId, error := strconv.ParseUint(parameters["userId"], 10, 32)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if followedUserId == followerId {
		responses.Error(w, http.StatusForbidden, errors.New("forbidden access"))
		return
	}

	db, error := database.Connection()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if error = repository.Follow(followedUserId, followerId); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, error := authentication.ExtractUserId(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}
	parameters := mux.Vars(r)
	userId, error := strconv.ParseUint(parameters["userId"], 10, 32)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}
	if followerId == userId {
		responses.Error(w, http.StatusForbidden, errors.New("unable to perform this operation"))
		return
	}
	db, error := database.Connection()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if error = repository.Unfollow(userId, followerId); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	responses.JSON(w, http.StatusOK, nil)
}
