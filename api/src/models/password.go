package models

type Password struct {
	New             string `json:"new"`
	CurrentPassword string `json:"current"`
}
