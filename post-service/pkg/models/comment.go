package models

import "time"

type Comment struct {
	CommentID string
	PostID    string
	UserID    string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
