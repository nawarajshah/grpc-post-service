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
		INSERT INTO users (user_id, email, password_hash, is_email_verified, verification_code, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query, user.UserID, user.Email, user.PasswordHash, user.IsEmailVerified, user.VerificationCode, user.CreatedAt)
	if err != nil {
		return fmt.Errorf("error inserting user: %w", err)
	}
	return nil
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	query := `
		SELECT user_id, email, password_hash, is_email_verified, created_at
		FROM users
		WHERE email = ?
	`
	row := r.db.QueryRow(query, email)

	var user models.User
	err := row.Scan(&user.UserID, &user.Email, &user.PasswordHash, &user.IsEmailVerified, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("No user found in GetByEmail with email: %s\n", email)
			return nil, nil // no user found
		}
		fmt.Printf("Error scanning row in GetByEmail: %v\n", err)
		return nil, fmt.Errorf("error retrieving user: %w", err)
	}

	fmt.Printf("User retrieved in GetByEmail: %+v\n", user)

	return &user, nil
}
