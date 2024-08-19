package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/models"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/repo"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/utils"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
	UserRepo         repo.UserRepository
	VerificationRepo repo.VerificationRepository
}

func NewAuthServiceServer(userRepo repo.UserRepository, verificationRepo repo.VerificationRepository) *AuthServiceServer {
	return &AuthServiceServer{
		UserRepo:         userRepo,
		VerificationRepo: verificationRepo,
	}
}

func (s *AuthServiceServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	// Validate email uniqueness
	existingUser, err := s.UserRepo.GetByEmail(req.User.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already in use")
	}

	// Hash the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Generate verification code
	verificationCode := utils.GenerateVerificationCode()

	// Create a new user
	user := &models.User{
		UserID:           utils.GenerateUUID(),
		Email:            req.User.Email,
		PasswordHash:     string(passwordHash),
		IsEmailVerified:  false,            // Email is not verified yet
		VerificationCode: verificationCode, // Store the verification code
		CreatedAt:        time.Now().Unix(),
	}

	err = s.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}

	// Send verification email
	go s.sendVerificationEmail(user.Email, verificationCode)

	return &pb.SignUpResponse{UserId: user.UserID}, nil
}

func (s *AuthServiceServer) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	// Log the request received by the gRPC service
	fmt.Printf("AuthServiceServer received SignIn request with email: %s, password: %s\n", req.Email, req.Password)

	// Get the user by email
	user, err := s.UserRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		fmt.Println("No user found with email:", req.Email)
		return nil, errors.New("invalid email or password")
	}

	// Check if the email is verified
	if !user.IsEmailVerified {
		return nil, errors.New("email not verified")
	}

	// Compare password hashes
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		fmt.Println("Password comparison failed:", err)
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.UserID)
	if err != nil {
		return nil, err
	}

	return &pb.SignInResponse{Token: token}, nil
}

func (s *AuthServiceServer) sendVerificationEmail(email, code string) {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_USERNAME"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Email Verification")
	m.SetBody("text/plain", fmt.Sprintf("Your verification code is: %s", code))

	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))

	d := gomail.NewDialer(os.Getenv("MAIL_HOST"), port, os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Failed to send email:", err)
	} else {
		fmt.Println("Verification email sent to:", email)
	}
}
