package service

import (
	"context"
	"fmt"
	"time"

	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/models"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/repo"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/emptypb" // Import the emptypb package
)

type PostServiceServer struct {
	pb.UnimplementedPostServiceServer
	Repo repo.PostRepository
}

func NewPostServiceServer(repo repo.PostRepository) *PostServiceServer {
	return &PostServiceServer{
		Repo: repo,
	}
}

func (s *PostServiceServer) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.PostResponse, error) {
	post := req.GetPost()

	// Validation
	if post.PostId == "" || post.Title == "" || post.Description == "" {
		return nil, fmt.Errorf("postId, title, and description are required")
	}
	if len(post.Title) > 100 {
		return nil, fmt.Errorf("title cannot exceed 100 characters")
	}

	createdAt := time.Now()
	updatedAt := time.Now()

	newPost := &models.Post{
		PostID:      post.PostId,
		Title:       post.Title,
		Description: post.Description,
		CreatedBy:   post.CreatedBy,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	err := s.Repo.Create(newPost)
	if err != nil {
		return nil, err
	}

	return &pb.PostResponse{
		Post: &pb.Post{
			PostId:      newPost.PostID,
			Title:       newPost.Title,
			Description: newPost.Description,
			CreatedBy:   newPost.CreatedBy,
			CreatedAt:   timestamppb.New(newPost.CreatedAt),
			UpdatedAt:   timestamppb.New(newPost.UpdatedAt),
		},
	}, nil
}

func (s *PostServiceServer) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.PostResponse, error) {
	postID := req.GetPostId()

	post, err := s.Repo.GetByID(postID)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, fmt.Errorf("post not found")
	}

	return &pb.PostResponse{
		Post: &pb.Post{
			PostId:      post.PostID,
			Title:       post.Title,
			Description: post.Description,
			CreatedBy:   post.CreatedBy,
			CreatedAt:   timestamppb.New(post.CreatedAt),
			UpdatedAt:   timestamppb.New(post.UpdatedAt),
		},
	}, nil
}

func (s *PostServiceServer) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.PostResponse, error) {
	post := req.GetPost()

	// Validation
	if post.PostId == "" || post.Title == "" || post.Description == "" {
		return nil, fmt.Errorf("postId, title, and description are required")
	}
	if len(post.Title) > 100 {
		return nil, fmt.Errorf("title cannot exceed 100 characters")
	}

	existingPost, err := s.Repo.GetByID(post.PostId)
	if err != nil {
		return nil, err
	}
	if existingPost == nil {
		return nil, fmt.Errorf("post not found")
	}

	existingPost.Title = post.Title
	existingPost.Description = post.Description
	existingPost.UpdatedAt = time.Now()

	err = s.Repo.Update(existingPost)
	if err != nil {
		return nil, err
	}

	return &pb.PostResponse{
		Post: &pb.Post{
			PostId:      existingPost.PostID,
			Title:       existingPost.Title,
			Description: existingPost.Description,
			CreatedBy:   existingPost.CreatedBy,
			CreatedAt:   timestamppb.New(existingPost.CreatedAt),
			UpdatedAt:   timestamppb.New(existingPost.UpdatedAt),
		},
	}, nil
}

func (s *PostServiceServer) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*emptypb.Empty, error) {
	postId := req.GetPostId()

	err := s.Repo.Delete(postId)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
