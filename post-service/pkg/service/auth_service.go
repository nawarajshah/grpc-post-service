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
	existingUser, err := s.UserRepo.GetByEmail(req.User.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already in use")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		UserID:          utils.GenerateUUID(),
		Email:           req.User.Email,
		PasswordHash:    string(passwordHash),
		IsEmailVerified: true,
	}

	err = s.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return &pb.SignUpResponse{UserId: user.UserID}, nil
}

func (s *AuthServiceServer) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	user, err := s.UserRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.UserID)
	if err != nil {
		return nil, err
	}

	return &pb.SignInResponse{Token: token}, nil
}
