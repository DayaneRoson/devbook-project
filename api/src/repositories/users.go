package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

// Find brings all users who fits the filter applied by name or nick
func (repository users) Find(name string) ([]models.User, error) {
	name = fmt.Sprintf("%%%s%%", name) //%name%

	rows, error := repository.db.Query(
		"select id, name, nick, email, createdAt from users where name like ? or nick like ?",
		name, name)
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if error = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); error != nil {
			return nil, error
		}

		users = append(users, user)
	}
	return users, nil
}

// FindById brings an specific user with a given Id
func (repository users) FindById(userId uint64) (models.User, error) {
	rows, error := repository.db.Query("select id, name, nick, email, createdAt from users where id = ?", userId)
	if error != nil {
		return models.User{}, error
	}
	defer rows.Close()
	var user models.User
	if rows.Next() {
		if error = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); error != nil {
			return models.User{}, error
		}
	}
	return user, nil
}
