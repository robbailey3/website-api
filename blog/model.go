package blog

import "time"

type Post struct {
	Id           string    `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	DateAdded    time.Time `json:"dateAdded"`
	DateModified time.Time `json:"dateModified"`
}
