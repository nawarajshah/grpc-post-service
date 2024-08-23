package models

import "time"

type Comment struct {
	CommentID  string
	PostID     string
	UserID     string
	Content    string
	CreatedAt  int64 // Store as Unix timestamp
	UpdatedAt  int64 // Store as Unix timestamp
	IsApproved bool
	OwnerID    string
}

// Convert time.Time to int64
func (c *Comment) SetCreatedAt(t time.Time) {
	c.CreatedAt = t.Unix()
}

func (c *Comment) SetUpdatedAt(t time.Time) {
	c.UpdatedAt = t.Unix()
}

// Convert int64 to time.Time when needed
func (c *Comment) GetCreatedAt() time.Time {
	return time.Unix(c.CreatedAt, 0)
}

func (c *Comment) GetUpdatedAt() time.Time {
	return time.Unix(c.UpdatedAt, 0)
}
