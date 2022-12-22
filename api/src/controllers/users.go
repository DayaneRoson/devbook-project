package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CreateUser creates an user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, error := io.ReadAll(r.Body)
	if error != nil {
		log.Fatal(error)
	}
	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		log.Fatal(error)
	}

	db, error := database.Connection()
	if error != nil {
		log.Fatal(error)
	}

	repository := repositories.NewUserRepository(db)
	userId, error := repository.Create(user)
	if error != nil {
		log.Fatal(error)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Query executed sucessfully. Id inserted: %d", userId)))
}

// GetUsers fetches all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Reading users"))
}

// GetUser fetches an user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Reading user"))
}

// UpdateUser updates an user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

// DeleteUser deletes an user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
