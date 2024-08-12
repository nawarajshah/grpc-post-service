package service

import (
	"context"

	"github.com/nawarajshah/grpc-post-service/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CommentService interface {
	CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CommentResponse, error)
	GetComment(ctx context.Context, req *pb.GetCommentRequest) (*pb.CommentResponse, error)
	UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.CommentResponse, error)
	DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*emptypb.Empty, error)
	ListComments(ctx context.Context, req *pb.ListCommentsRequest) (*pb.ListCommentsResponse, error)
}

type commentService struct {
	client pb.CommentServiceClient
}

func NewCommentService(client pb.CommentServiceClient) CommentService {
	return &commentService{client: client}
}

func (s *commentService) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CommentResponse, error) {
	return s.client.CreateComment(ctx, req)
}

func (s *commentService) GetComment(ctx context.Context, req *pb.GetCommentRequest) (*pb.CommentResponse, error) {
	return s.client.GetComment(ctx, req)
}

func (s *commentService) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.CommentResponse, error) {
	return s.client.UpdateComment(ctx, req)
}

func (s *commentService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*emptypb.Empty, error) {
	return s.client.DeleteComment(ctx, req)
}

func (s *commentService) ListComments(ctx context.Context, req *pb.ListCommentsRequest) (*pb.ListCommentsResponse, error) {
	return s.client.ListComments(ctx, req)
}
