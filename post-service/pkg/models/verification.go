package models

type Verification struct {
	UserID           string
	VerificationCode string
	CreatedAt        int64
}
