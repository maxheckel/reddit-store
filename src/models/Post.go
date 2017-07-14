package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Text string
	Link string
	Comments []Comment
	Hash string
	Subreddit Subreddit
}
