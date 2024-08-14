package repo

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/nawarajshah/grpc-post-service/post-service/pkg/models"
)

type PostRepository interface {
	Create(post *models.Post) error
	GetByID(postID string) (*models.Post, error)
	Update(post *models.Post) error
	Delete(postID string) error
}

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(post *models.Post) error {
	query := "INSERT INTO posts (postid, title, description, created_by, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, post.PostID, post.Title, post.Description, post.CreatedBy, post.CreatedAt.Unix(), post.UpdatedAt.Unix())
	if err != nil {
		return fmt.Errorf("error inserting post: %w", err)
	}
	return nil
}

func (r *postRepository) GetByID(postID string) (*models.Post, error) {
	query := "SELECT postid, title, description, created_by, created_at, updated_at FROM posts WHERE postid = ?"
	row := r.db.QueryRow(query, postID)

	var post models.Post
	var createdAtUnix, updatedAtUnix int64
	err := row.Scan(&post.PostID, &post.Title, &post.Description, &post.CreatedBy, &createdAtUnix, &updatedAtUnix)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no post found
		}
		return nil, fmt.Errorf("error retrieving post: %w", err)
	}

	post.CreatedAt = time.Unix(createdAtUnix, 0)
	post.UpdatedAt = time.Unix(updatedAtUnix, 0)

	return &post, nil
}

func (r *postRepository) Update(post *models.Post) error {
	query := "UPDATE posts SET title = ?, description = ?, updated_at = ? WHERE postid = ?"
	_, err := r.db.Exec(query, post.Title, post.Description, post.UpdatedAt.Unix(), post.PostID)
	if err != nil {
		return fmt.Errorf("error updating post: %w", err)
	}
	return nil
}

func (r *postRepository) Delete(postID string) error {
	query := "DELETE FROM posts WHERE postid = ?"
	result, err := r.db.Exec(query, postID)
	if err != nil {
		return fmt.Errorf("error deleting post: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("post not found")
	}

	return nil
}
