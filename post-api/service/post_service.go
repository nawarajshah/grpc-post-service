package service

import (
	"context"
	"fmt"

	"github.com/nawarajshah/grpc-post-service/pb"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PostService interface {
	CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.PostResponse, error)
	GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.PostResponse, error)
	UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.PostResponse, error)
	DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*emptypb.Empty, error)
}

type postService struct {
	client pb.PostServiceClient
}

func NewPostService(client pb.PostServiceClient) PostService {
	return &postService{client: client}
}

func (s *postService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.PostResponse, error) {
	res, err := s.client.CreatePost(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			return nil, fmt.Errorf("gRPC error: %v", st.Message())
		}
		return nil, fmt.Errorf("unexpected error: %v", err)
	}
	return res, nil
}

func (s *postService) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.PostResponse, error) {
	res, err := s.client.GetPost(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			return nil, fmt.Errorf("gRPC error: %v", st.Message())
		}
		return nil, fmt.Errorf("unexpected error: %v", err)
	}
	return res, nil
}

func (s *postService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.PostResponse, error) {
	res, err := s.client.UpdatePost(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			return nil, fmt.Errorf("gRPC error: %v", st.Message())
		}
		return nil, fmt.Errorf("unexpected error: %v", err)
	}
	return res, nil
}

func (s *postService) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*emptypb.Empty, error) {
	res, err := s.client.DeletePost(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			return nil, fmt.Errorf("gRPC error: %v", st.Message())
		}
		return nil, fmt.Errorf("unexpected error: %v", err)
	}
	return res, nil
}
