package models

type User struct {
	UserID           string
	Email            string
	PasswordHash     string
	IsEmailVerified  bool
	VerificationCode string
	CreatedAt        int64
}
