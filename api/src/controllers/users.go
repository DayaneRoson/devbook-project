package controllers

import "net/http"

//CreateUser creates an user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user"))
}

//GetUsers fetches all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Reading users"))
}

//GetUser fetches an user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Reading user"))
}

//UpdateUser updates an user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

//DeleteUser deletes an user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
