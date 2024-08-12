package service

import (
	"context"
	"fmt"

	"github.com/nawarajshah/grpc-post-service/pb"
	"google.golang.org/grpc/status"
)

type CommentService interface {
	CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CommentResponse, error)
	GetComment(ctx context.Context, req *pb.GetCommentRequest) (*pb.CommentResponse, error)
	UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.CommentResponse, error)
	DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.Empty, error)
	ListComments(ctx context.Context, req *pb.ListCommentsRequest) (*pb.ListCommentsResponse, error)
}

type commentService struct {
	client pb.CommentServiceClient
}

func NewCommentService(client pb.CommentServiceClient) CommentService {
	return &commentService{client: client}
}

func (s *commentService) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CommentResponse, error) {
	res, err := s.client.CreateComment(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			return nil, fmt.Errorf("gRPC error: %v", st.Message())
		}
		return nil, fmt.Errorf("unexpected error: %v", err)
	}
	return res, nil
}

func (s *commentService) GetComment(ctx context.Context, req *pb.GetCommentRequest) (*pb.CommentResponse, error) {
	res, err := s.client.GetComment(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			return nil, fmt.Errorf("gRPC error: %v", st.Message())
		}
		return nil, fmt.Errorf("unexpected error: %v", err)
	}
	return res, nil
}

func (s *commentService) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.CommentResponse, error) {
	res, err := s.client.UpdateComment(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			return nil, fmt.Errorf("gRPC error: %v", st.Message())
		}
		return nil, fmt.Errorf("unexpected error: %v", err)
	}
	return res, nil
}

func (s *commentService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.Empty, error) {
	res, err := s.client.DeleteComment(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			return nil, fmt.Errorf("gRPC error: %v", st.Message())
		}
		return nil, fmt.Errorf("unexpected error: %v", err)
	}
	return res, nil
}

func (s *commentService) ListComments(ctx context.Context, req *pb.ListCommentsRequest) (*pb.ListCommentsResponse, error) {
	res, err := s.client.ListComments(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			return nil, fmt.Errorf("gRPC error: %v", st.Message())
		}
		return nil, fmt.Errorf("unexpected error: %v", err)
	}
	return res, nil
}
