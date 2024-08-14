package service

import (
	"context"

	"github.com/nawarajshah/grpc-post-service/pb"
)

type AuthService interface {
	SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error)
	SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error)
}

type authService struct {
	client pb.AuthServiceClient
}

func NewAuthService(client pb.AuthServiceClient) AuthService {
	return &authService{client: client}
}

func (s *authService) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	return s.client.SignUp(ctx, req)
}

func (s *authService) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	return s.client.SignIn(ctx, req)
}
