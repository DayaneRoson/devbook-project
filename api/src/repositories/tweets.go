package repositories

import (
	"api/src/models"
	"database/sql"
)

type Tweets struct {
	db *sql.DB
}

func NewTweetRepository(db *sql.DB) *Tweets {
	return &Tweets{db}
}

// Create inserts a tweet in the database
func (repository Tweets) Create(tweet models.Tweet) (uint64, error) {
	statement, error := repository.db.Prepare("insert into tweets (title, content, author_id) values (?,?,?)")
	if error != nil {
		return 0, error
	}
	defer statement.Close()
	result, error := statement.Exec(tweet.Title, tweet.Content, tweet.AuthorId)
	if error != nil {
		return 0, error
	}

	lastInsertId, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastInsertId), nil
}
