package entity

import (
	"time"
)

type Post struct {
	ID           string
	Text         string
	AuthorUserID string
	CreatedAt    time.Time
}
