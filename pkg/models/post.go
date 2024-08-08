package models

import "time"

type Post struct {
	PostID      string
	Title       string
	Description string
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
