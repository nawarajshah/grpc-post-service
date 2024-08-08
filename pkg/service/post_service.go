package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/pkg/models"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type PostServiceServer struct {
	pb.UnimplementedPostServiceServer
	DB *sql.DB
}

func NewPostServiceServer(db *sql.DB) *PostServiceServer {
	return &PostServiceServer{
		DB: db,
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

	query := "INSERT INTO posts (postid, title, description, created_by, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := s.DB.Exec(query, post.PostId, post.Title, post.Description, post.CreatedBy, createdAt.Unix(), updatedAt.Unix())
	if err != nil {
		return nil, fmt.Errorf("error inserting post: %v", err)
	}

	return &pb.PostResponse{
		Post: &pb.Post{
			PostId:     post.PostId,
			Title:      post.Title,
			Description: post.Description,
			CreatedBy:  post.CreatedBy,
			CreatedAt:  timestamppb.New(createdAt),
			UpdatedAt:  timestamppb.New(updatedAt),
		},
	}, nil
}

func (s *PostServiceServer) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.PostResponse, error) {
	postId := req.GetPostId()

	query := "SELECT postid, title, description, created_by, created_at, updated_at FROM posts WHERE postid = ?"
	row := s.DB.QueryRow(query, postId)

	var post models.Post
	var createdAtUnix, updatedAtUnix int64
	err := row.Scan(&post.PostID, &post.Title, &post.Description, &post.CreatedBy, &createdAtUnix, &updatedAtUnix)
	if err != nil {
		return nil, fmt.Errorf("error retrieving post: %v", err)
	}

	post.CreatedAt = time.Unix(createdAtUnix, 0)
	post.UpdatedAt = time.Unix(updatedAtUnix, 0)

	return &pb.PostResponse{
		Post: &pb.Post{
			PostId:     post.PostID,
			Title:      post.Title,
			Description: post.Description,
			CreatedBy:  post.CreatedBy,
			CreatedAt:  timestamppb.New(post.CreatedAt),
			UpdatedAt:  timestamppb.New(post.UpdatedAt),
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

	updatedAt := time.Now()

	query := "UPDATE posts SET title = ?, description = ?, updated_at = ? WHERE postid = ?"
	_, err := s.DB.Exec(query, post.Title, post.Description, updatedAt.Unix(), post.PostId)
	if err != nil {
		return nil, fmt.Errorf("error updating post: %v", err)
	}

	return &pb.PostResponse{
		Post: &pb.Post{
			PostId:     post.PostId,
			Title:      post.Title,
			Description: post.Description,
			CreatedBy:  post.CreatedBy,
			CreatedAt:  post.CreatedAt,
			UpdatedAt:  timestamppb.New(updatedAt),
		},
	}, nil
}

func (s *PostServiceServer) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.Empty, error) {
	postId := req.GetPostId()

	query := "DELETE FROM posts WHERE postid = ?"
	_, err := s.DB.Exec(query, postId)
	if err != nil {
		return nil, fmt.Errorf("error deleting post: %v", err)
	}

	return &pb.Empty{}, nil
}
