package service

import (
	"context"
	"fmt"
	"time"

	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/models"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/repo"
)

type GrpcService struct {
	pb.UnimplementedPostServiceServer
	PostRepo repo.PostRepository
}

// NewPostServiceServer is a constructor for GrpcService
func NewPostServiceServer(postRepo repo.PostRepository) *GrpcService {
	return &GrpcService{
		PostRepo: postRepo,
	}
}

func (s *GrpcService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.PostResponse, error) {
	post := &models.Post{
		PostID:      req.Post.PostId,
		Title:       req.Post.Title,
		Description: req.Post.Description,
		UserID:      req.Post.UserId,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	err := s.PostRepo.Create(post)
	if err != nil {
		return nil, fmt.Errorf("error creating post: %v", err)
	}

	return &pb.PostResponse{
		Post: &pb.Post{
			PostId:      post.PostID,
			Title:       post.Title,
			Description: post.Description,
			UserId:      post.UserID,
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
		},
	}, nil
}

func (s *GrpcService) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.PostResponse, error) {
	post, err := s.PostRepo.GetByID(req.PostId)
	if err != nil {
		return nil, fmt.Errorf("error retrieving post: %v", err)
	}
	if post == nil {
		return nil, fmt.Errorf("post not found")
	}

	return &pb.PostResponse{
		Post: &pb.Post{
			PostId:      post.PostID,
			Title:       post.Title,
			Description: post.Description,
			UserId:      post.UserID,
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
		},
	}, nil
}

func (s *GrpcService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.PostResponse, error) {
	post, err := s.PostRepo.GetByID(req.Post.PostId)
	if err != nil {
		return nil, fmt.Errorf("error retrieving post: %v", err)
	}
	if post == nil {
		return nil, fmt.Errorf("post not found")
	}

	post.Title = req.Post.Title
	post.Description = req.Post.Description
	post.UpdatedAt = time.Now().Unix()

	err = s.PostRepo.Update(post)
	if err != nil {
		return nil, fmt.Errorf("error updating post: %v", err)
	}

	return &pb.PostResponse{
		Post: &pb.Post{
			PostId:      post.PostID,
			Title:       post.Title,
			Description: post.Description,
			UserId:      post.UserID,
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
		},
	}, nil
}

func (s *GrpcService) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	err := s.PostRepo.Delete(req.PostId)
	if err != nil {
		return nil, fmt.Errorf("error deleting post: %v", err)
	}

	return &pb.DeletePostResponse{
		Status: "Post deleted successfully",
	}, nil
}
