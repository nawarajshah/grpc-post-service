package service

import (
	"context"

	"github.com/nawarajshah/grpc-post-service/pb"
)

type CommentService interface {
	CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CommentResponse, error)
	GetComment(ctx context.Context, req *pb.GetCommentRequest) (*pb.CommentResponse, error)
	DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error)
	ListComments(ctx context.Context, req *pb.GetCommentsByPostIDRequest) (*pb.GetCommentsByPostIDResponse, error)
	ApproveComment(ctx context.Context, req *pb.ApproveCommentRequest) (*pb.CommentResponse, error)
	UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.CommentResponse, error) // Add this line
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
	return s.client.GetCommentByID(ctx, req) // Use the correct method name
}

func (s *commentService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	return s.client.DeleteComment(ctx, req)
}

func (s *commentService) ListComments(ctx context.Context, req *pb.GetCommentsByPostIDRequest) (*pb.GetCommentsByPostIDResponse, error) {
	return s.client.GetCommentsByPostID(ctx, req) // Use the correct method name
}

func (s *commentService) ApproveComment(ctx context.Context, req *pb.ApproveCommentRequest) (*pb.CommentResponse, error) {
	return s.client.ApproveComment(ctx, req)
}

func (s *commentService) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.CommentResponse, error) { // Implement this method
	return s.client.UpdateComment(ctx, req)
}
