package service

import (
	"context"

	"github.com/nawarajshah/grpc-post-service/pb"
)

type VerificationService interface {
	VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error)
}

type verificationService struct {
	client pb.VerificationServiceClient
}

func NewVerificationService(client pb.VerificationServiceClient) VerificationService {
	return &verificationService{client: client}
}

func (s *verificationService) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error) {
	return s.client.VerifyEmail(ctx, req)
}
