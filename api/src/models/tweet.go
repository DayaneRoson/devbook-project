package models

import (
	"errors"
	"strings"
	"time"
)

type Tweet struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorId   uint64    `json:"authorId,omitempty"`
	AuthorNIck string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

func (tweet *Tweet) Prepare() error {
	if error := tweet.validate(); error != nil {
		return error
	}

	tweet.format()
	return nil
}

func (tweet *Tweet) validate() error {
	if tweet.Title == "" || tweet.Title == " " {
		return errors.New("title is required")
	}
	if tweet.Content == "" || tweet.Content == " " {
		return errors.New("content is required")
	}
	return nil
}

func (tweet *Tweet) format() {
	tweet.Title = strings.TrimSpace(tweet.Title)
	tweet.Content = strings.TrimSpace(tweet.Content)
}
