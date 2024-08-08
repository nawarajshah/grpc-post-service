package server

import (
	"context"
	"database/sql"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nawarajshah/grpc-post-service/postpb"
)

type Server struct {
	postpb.UnimplementedPostServiceServer
	DB *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{DB: db}
}

// CreatePost creates a new post
func (s *Server) CreatePost(ctx context.Context, req *postpb.CreatePostRequest) (*postpb.CreatePostResponse, error) {
	post := req.GetPost()

	if post.PostId == "" || post.Title == "" || post.Description == "" {
		return nil, status.Errorf(codes.InvalidArgument, "post_id, title, and description are required")
	}

	if len(post.Title) > 100 {
		return nil, status.Errorf(codes.InvalidArgument, "title cannot exceed 100 characters")
	}

	createdAt := time.Now().Unix()
	query := "INSERT INTO posts (postid, title, description, created_by, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := s.DB.ExecContext(ctx, query, post.PostId, post.Title, post.Description, post.CreatedBy, createdAt, createdAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert post: %v", err)
	}

	return &postpb.CreatePostResponse{PostId: post.PostId}, nil
}

// ReadPost fetches a post by ID
func (s *Server) ReadPost(ctx context.Context, req *postpb.ReadPostRequest) (*postpb.ReadPostResponse, error) {
	postID := req.GetPostId()

	query := "SELECT postid, title, description, created_by, created_at, updated_at FROM posts WHERE postid = ?"
	row := s.DB.QueryRowContext(ctx, query, postID)

	post := &postpb.Post{}
	err := row.Scan(&post.PostId, &post.Title, &post.Description, &post.CreatedBy, &post.CreatedAt, &post.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, status.Errorf(codes.NotFound, "post not found")
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch post: %v", err)
	}

	return &postpb.ReadPostResponse{Post: post}, nil
}

// UpdatePost updates an existing post
func (s *Server) UpdatePost(ctx context.Context, req *postpb.UpdatePostRequest) (*postpb.UpdatePostResponse, error) {
	post := req.GetPost()

	if post.PostId == "" || post.Title == "" || post.Description == "" {
		return nil, status.Errorf(codes.InvalidArgument, "post_id, title, and description are required")
	}

	if len(post.Title) > 100 {
		return nil, status.Errorf(codes.InvalidArgument, "title cannot exceed 100 characters")
	}

	updatedAt := time.Now().Unix()
	query := "UPDATE posts SET title = ?, description = ?, updated_at = ? WHERE postid = ?"
	res, err := s.DB.ExecContext(ctx, query, post.Title, post.Description, updatedAt, post.PostId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update post: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "post not found")
	}

	return &postpb.UpdatePostResponse{Success: true}, nil
}

// DeletePost deletes a post by ID
func (s *Server) DeletePost(ctx context.Context, req *postpb.DeletePostRequest) (*postpb.DeletePostResponse, error) {
	postID := req.GetPostId()

	query := "DELETE FROM posts WHERE postid = ?"
	res, err := s.DB.ExecContext(ctx, query, postID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete post: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "post not found")
	}

	return &postpb.DeletePostResponse{Success: true}, nil
}
