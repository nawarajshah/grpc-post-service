package models

type Post struct {
	PostID      string
	Title       string
	Description string
	UserID      string // This should match 'created_by' in your database
	CreatedAt   int64
	UpdatedAt   int64
}
