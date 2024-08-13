package models

type User struct {
	UserID          string
	Email           string
	PasswordHash    string
	IsEmailVerified bool
}
