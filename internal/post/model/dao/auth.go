package dao

import (
	"time"
)

type Post struct {
	PostID          string    `db:"post_id"`
	Title           string    `db:"title"`
	Content         string    `db:"content"`
	Author          string    `db:"author"`
	PublicationDate time.Time `db:"publication_date"`
	Tags            string    `db:"tags"`
}
