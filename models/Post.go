package models


type Post struct {
	Id string
	Title string
	Url string
	Comments []Comment
	Score int
	SubredditID uint
	Created float32
}
