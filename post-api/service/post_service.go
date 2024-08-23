package service

import (
	"context"
	"github.com/nawarajshah/grpc-post-service/pb"
)

type PostService interface {
	CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.PostResponse, error)
	GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.PostResponse, error)
	UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.PostResponse, error)
	DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) // Update method signature
}

type postService struct {
	client pb.PostServiceClient
}

func NewPostService(client pb.PostServiceClient) PostService {
	return &postService{client: client}
}

func (s *postService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.PostResponse, error) {
	return s.client.CreatePost(ctx, req)
}

func (s *postService) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.PostResponse, error) {
	return s.client.GetPost(ctx, req)
}

func (s *postService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.PostResponse, error) {
	return s.client.UpdatePost(ctx, req)
}

func (s *postService) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	return s.client.DeletePost(ctx, req)
}
