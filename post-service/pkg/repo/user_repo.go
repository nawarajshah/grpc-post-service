package repo

import (
	"database/sql"
	"fmt"

	"github.com/nawarajshah/grpc-post-service/post-service/pkg/models"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (user_id, email, password_hash, is_email_verified)
		VALUES (?, ?, ?, ?)
	`
	_, err := r.db.Exec(query, user.UserID, user.Email, user.PasswordHash, user.IsEmailVerified)
	if err != nil {
		return fmt.Errorf("error inserting user: %w", err)
	}
	return nil
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	query := `
		SELECT user_id, email, password_hash, is_email_verified
		FROM users
		WHERE email = ?
	`
	row := r.db.QueryRow(query, email)

	var user models.User
	err := row.Scan(&user.UserID, &user.Email, &user.PasswordHash, &user.IsEmailVerified)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no user found
		}
		return nil, fmt.Errorf("error retrieving user: %w", err)
	}

	return &user, nil
}
