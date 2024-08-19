package service

import (
	"context"

	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/repo"
)

type VerificationServiceServer struct {
	pb.UnimplementedVerificationServiceServer
	VerificationRepo repo.VerificationRepository
}

func NewVerificationServiceServer(verificationRepo repo.VerificationRepository) *VerificationServiceServer {
	return &VerificationServiceServer{
		VerificationRepo: verificationRepo,
	}
}

func (s *VerificationServiceServer) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error) {
	err := s.VerificationRepo.VerifyEmail(req.Email, req.VerificationCode)
	if err != nil {
		return nil, err
	}
	return &pb.VerifyEmailResponse{Message: "Email verified successfully"}, nil
}
