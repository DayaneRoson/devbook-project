package repositories

import (
	"api/src/models"
	"database/sql"
)

type tweets struct {
	db *sql.DB
}

func NewTweetRepository(db *sql.DB) *tweets {
	return &tweets{db}
}

// Create inserts a tweets in the database
func (repository tweets) Create(tweet models.Tweet) (uint64, error) {
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

// FindById brings an specific tweets from the database
func (repository tweets) FindById(tweetId uint64) (models.Tweet, error) {
	rows, error := repository.db.Query(
		`select t.*, u.nick from 
		tweets t inner join users u
		on u.id = t.author_id where t.id = ?`, tweetId)
	if error != nil {
		return models.Tweet{}, error
	}
	defer rows.Close()

	var tweet models.Tweet
	if rows.Next() {
		if error = rows.Scan(&tweet.ID, &tweet.Title, &tweet.Content,
			&tweet.AuthorId, &tweet.Likes, &tweet.CreatedAt, &tweet.AuthorNIck); error != nil {
			return models.Tweet{}, error
		}
	}
	return tweet, nil
}

func (repository tweets) Find(userId uint64) ([]models.Tweet, error) {
	rows, error := repository.db.Query(`
	select distinct t.*, u.nick from tweets t
	inner join users u on u.id = t.author_id
	inner join followers f on t.author_id = f.user_id
	where u.id = ? or f.follower_id = ?
	order by 1 desc`, userId, userId)
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	var tweets []models.Tweet
	for rows.Next() {
		var tweet models.Tweet

		if error = rows.Scan(&tweet.ID, &tweet.Title, &tweet.Content,
			&tweet.AuthorId, &tweet.Likes, &tweet.CreatedAt, &tweet.AuthorNIck); error != nil {
			return nil, error
		}

		tweets = append(tweets, tweet)
	}

	return tweets, nil
}

func (repository tweets) Update(tweetId uint64, tweet models.Tweet) error {
	statement, error := repository.db.Prepare("update tweets set title = ?, content = ? where author_id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(tweet.Title, tweet.Content, tweetId); error != nil {
		return error
	}

	return nil
}
