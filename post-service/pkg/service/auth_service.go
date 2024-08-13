package service

import (
	"context"
	"errors"

	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/models"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/repo"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
	UserRepo repo.UserRepository
}

func NewAuthServiceServer(userRepo repo.UserRepository) *AuthServiceServer {
	return &AuthServiceServer{
		UserRepo: userRepo,
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

	// Create a new user
	user := &models.User{
		UserID:          utils.GenerateUUID(),
		Email:           req.User.Email,
		PasswordHash:    string(passwordHash),
		IsEmailVerified: true, // for now
	}

	err = s.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return &pb.SignUpResponse{UserId: user.UserID}, nil
}

func (s *AuthServiceServer) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	// Get the user by email
	user, err := s.UserRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	// Compare password hashes
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.UserID)
	if err != nil {
		return nil, err
	}

	return &pb.SignInResponse{Token: token}, nil
}
