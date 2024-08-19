package repo

import (
	"database/sql"
	"fmt"
)

type VerificationRepository interface {
	SetVerificationCode(userID, code string) error
	VerifyEmail(email, code string) error
}

type verificationRepository struct {
	db *sql.DB
}

func NewVerificationRepository(db *sql.DB) VerificationRepository {
	return &verificationRepository{db: db}
}

func (r *verificationRepository) SetVerificationCode(userID, code string) error {
	query := `
		UPDATE users
		SET verification_code = ?
		WHERE user_id = ?
	`
	_, err := r.db.Exec(query, code, userID)
	if err != nil {
		return fmt.Errorf("error setting verification code: %w", err)
	}
	return nil
}

func (r *verificationRepository) VerifyEmail(email, code string) error {
	query := `
		UPDATE users
		SET is_email_verified = true, verification_code = NULL
		WHERE email = ? AND verification_code = ?
	`
	result, err := r.db.Exec(query, email, code)
	if err != nil {
		return fmt.Errorf("error verifying email: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return fmt.Errorf("invalid verification code or email")
	}
	return nil
}
