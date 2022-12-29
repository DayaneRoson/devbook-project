package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
)

// Login authenticates an user in the API
func Login(w http.ResponseWriter, r *http.Request) {
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

	db, error := database.Connection()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	userSaved, error := repository.FindByEmail(user.Email)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if error = security.VerifyPassword(userSaved.Password, user.Password); error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	token, _ := authentication.CreateToken(uint64(userSaved.ID))
	w.Write([]byte(token))
}
