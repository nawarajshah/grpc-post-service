package repo

import (
	"database/sql"
	"fmt"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/models"
	"log"
)

type CommentRepository interface {
	Create(comment *models.Comment) error
	GetByID(commentID string) (*models.Comment, error)
	GetByPostID(postID string) ([]*models.Comment, error)
	Update(comment *models.Comment) error
	ApproveComment(commentID string) error
	Delete(commentID string) error
}

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(comment *models.Comment) error {
	query := `
		INSERT INTO comments (commentid, postid, userid, content, created_at, updated_at, is_approved, owner_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query, comment.CommentID, comment.PostID, comment.UserID, comment.Content, comment.CreatedAt, comment.UpdatedAt, comment.IsApproved, comment.OwnerID)
	if err != nil {
		return fmt.Errorf("error inserting comment: %w", err)
	}
	return nil
}

func (r *commentRepository) GetByID(commentID string) (*models.Comment, error) {
	query := `
		SELECT commentid, postid, userid, content, created_at, updated_at, is_approved, owner_id
		FROM comments
		WHERE commentid = ?
	`
	row := r.db.QueryRow(query, commentID)

	var comment models.Comment
	var createdAtUnix, updatedAtUnix int64
	err := row.Scan(&comment.CommentID, &comment.PostID, &comment.UserID, &comment.Content, &createdAtUnix, &updatedAtUnix, &comment.IsApproved, &comment.OwnerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no comment found
		}
		return nil, fmt.Errorf("error retrieving comment: %w", err)
	}

	comment.CreatedAt = createdAtUnix
	comment.UpdatedAt = updatedAtUnix

	return &comment, nil
}

func (r *commentRepository) GetByPostID(postID string) ([]*models.Comment, error) {
	query := `
		SELECT commentid, postid, userid, content, created_at, updated_at, is_approved, owner_id
		FROM comments
		WHERE postid = ? AND is_approved = true
	`
	rows, err := r.db.Query(query, postID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving comments: %w", err)
	}
	defer rows.Close()

	var comments []*models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.CommentID, &comment.PostID, &comment.UserID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt, &comment.IsApproved, &comment.OwnerID)
		if err != nil {
			return nil, fmt.Errorf("error scanning comment: %w", err)
		}
		comments = append(comments, &comment)
	}

	return comments, nil
}

func (r *commentRepository) Update(comment *models.Comment) error {
	query := `
		UPDATE comments
		SET content = ?, updated_at = ?, is_approved = ?
		WHERE commentid = ?
	`
	_, err := r.db.Exec(query, comment.Content, comment.UpdatedAt, comment.IsApproved, comment.CommentID)
	if err != nil {
		return fmt.Errorf("error updating comment: %w", err)
	}
	return nil
}

func (r *commentRepository) ApproveComment(commentID string) error {
	query := `
		UPDATE comments
		SET is_approved = true
		WHERE commentid = ?
	`
	_, err := r.db.Exec(query, commentID)
	if err != nil {
		return fmt.Errorf("error approving comment: %w", err)
	}
	return nil
}

func (r *commentRepository) Delete(commentID string) error {
	query := `
		DELETE FROM comments
		WHERE commentid = ?
	`
	_, err := r.db.Exec(query, commentID)
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
		// Directly assign the Unix timestamp to int64 fields
		comment.CreatedAt = createdAtUnix
		comment.UpdatedAt = updatedAtUnix
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
