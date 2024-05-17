package dto

import "time"

type PostCreateRequest struct {
}

type Post struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Content         string    `json:"content"`
	Author          string    `json:"author"`
	PublicationDate time.Time `json:"publication_date"`
	Tags            string    `json:"tags"`
}
