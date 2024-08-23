package models

type Post struct {
	PostID      string `db:"postid"`
	Title       string `db:"title"`
	Description string `db:"description"`
	CreatedBy   string `db:"created_by"` // Update this field name
	CreatedAt   int64  `db:"created_at"`
	UpdatedAt   int64  `db:"updated_at"`
}
