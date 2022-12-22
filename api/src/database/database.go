package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Connection opens the connection to the database
func Connection() (*sql.DB, error) {
	db, error := sql.Open("mysql", config.StringDatabaseConnection)
	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		db.Close()
		return nil, error
	}

	return db, nil
}
