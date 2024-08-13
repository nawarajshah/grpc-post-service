package repo

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/nawarajshah/grpc-post-service/post-service/pkg/models"
)

type CommentRepository interface {
	Create(comment *models.Comment) error
	GetByID(postID, commentID string) (*models.Comment, error)
	Update(comment *models.Comment) error
	Delete(postID, commentID string) error
	ListByPostID(postID string) ([]*models.Comment, error)
}

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(comment *models.Comment) error {
	// Check if postId exists
	if !r.postExists(comment.PostID) {
		return fmt.Errorf("post with id %s does not exist", comment.PostID)
	}

	query := `
		INSERT INTO comments (commentid, postid, userid, content, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query, comment.CommentID, comment.PostID, comment.UserID, comment.Content, comment.CreatedAt.Unix(), comment.UpdatedAt.Unix())
	if err != nil {
		return fmt.Errorf("error inserting comment: %w", err)
	}
	return nil
}

func (r *commentRepository) GetByID(postID, commentID string) (*models.Comment, error) {
	// Check if postId exists
	if !r.postExists(postID) {
		return nil, fmt.Errorf("post with id %s does not exist", postID)
	}

	query := `
		SELECT commentid, postid, userid, content, created_at, updated_at
		FROM comments
		WHERE postid = ? AND commentid = ?
	`
	row := r.db.QueryRow(query, postID, commentID)

	var comment models.Comment
	var createdAtUnix, updatedAtUnix int64
	err := row.Scan(&comment.CommentID, &comment.PostID, &comment.UserID, &comment.Content, &createdAtUnix, &updatedAtUnix)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no comment found
		}
		return nil, fmt.Errorf("error retrieving comment: %w", err)
	}

	comment.CreatedAt = time.Unix(createdAtUnix, 0)
	comment.UpdatedAt = time.Unix(updatedAtUnix, 0)

	return &comment, nil
}

func (r *commentRepository) Update(comment *models.Comment) error {
	// Check if postId exists
	if !r.postExists(comment.PostID) {
		return fmt.Errorf("post with id %s does not exist", comment.PostID)
	}

	query := `
		UPDATE comments
		SET content = ?, updated_at = ?
		WHERE postid = ? AND commentid = ?
	`
	_, err := r.db.Exec(query, comment.Content, comment.UpdatedAt.Unix(), comment.PostID, comment.CommentID)
	if err != nil {
		return fmt.Errorf("error updating comment: %w", err)
	}
	return nil
}

func (r *commentRepository) Delete(postID, commentID string) error {
	// Check if postId exists
	if !r.postExists(postID) {
		return fmt.Errorf("post with id %s does not exist", postID)
	}

	query := `
		DELETE FROM comments
		WHERE postid = ? AND commentid = ?
	`
	_, err := r.db.Exec(query, postID, commentID)
	if err != nil {
		return fmt.Errorf("error deleting comment: %w", err)
	}
	return nil
}

func (r *commentRepository) ListByPostID(postID string) ([]*models.Comment, error) {
	// Check if postId exists
	if !r.postExists(postID) {
		return nil, fmt.Errorf("post with id %s does not exist", postID)
	}

	query := `
		SELECT commentid, postid, userid, content, created_at, updated_at
		FROM comments
		WHERE postid = ?
	`
	rows, err := r.db.Query(query, postID)
	if err != nil {
		return nil, fmt.Errorf("error listing comments: %w", err)
	}
	defer rows.Close()

	var comments []*models.Comment
	for rows.Next() {
		var comment models.Comment
		var createdAtUnix, updatedAtUnix int64
		if err := rows.Scan(&comment.CommentID, &comment.PostID, &comment.UserID, &comment.Content, &createdAtUnix, &updatedAtUnix); err != nil {
			return nil, fmt.Errorf("error scanning comment: %w", err)
		}
		comment.CreatedAt = time.Unix(createdAtUnix, 0)
		comment.UpdatedAt = time.Unix(updatedAtUnix, 0)
		comments = append(comments, &comment)
	}
	return comments, nil
}

// Helper method to check if a postId exists
func (r *commentRepository) postExists(postID string) bool {
	var exists bool
	query := `
		SELECT EXISTS (
			SELECT 1 FROM posts WHERE postid = ?
		)
	`
	err := r.db.QueryRow(query, postID).Scan(&exists)
	if err != nil {
		log.Printf("Error checking if post exists: %v", err)
		return false
	}
	return exists
}
