package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

// NewUserRepository creates a user repository
func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

// Create inserts a new user in the database
func (repository users) Create(user models.User) (uint32, error) {
	statement, error := repository.db.Prepare(
		"insert into users (name, nick, email, password) values (?,?,?,?)")
	if error != nil {
		return 0, error
	}
	defer statement.Close()
	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if error != nil {
		return 0, nil
	}

	lastInsertId, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint32(lastInsertId), nil
}
