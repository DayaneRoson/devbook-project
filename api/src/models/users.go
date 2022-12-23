package models

import (
	"errors"
	"strings"
	"time"
)

// user represents an user using the social media
type User struct {
	ID        uint32    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare calls the methods to validate e format the user received
func (user *User) Prepare() error {
	if error := user.validate(); error != nil {
		return error
	}
	user.format()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" || user.Name == " " {
		return errors.New("field name is required")
	}
	if user.Nick == "" || user.Nick == " " {
		return errors.New("field NIck is required")
	}
	if user.Email == "" || user.Email == " " {
		return errors.New("field email is required")
	}
	if user.Password == "" {
		return errors.New("field password is required")
	}
	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
