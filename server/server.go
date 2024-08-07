package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/nawarajshah/grpc-post-service/pb"
)

const (
	port       = ":50051"
	dsn        = "nawaraj:nawaraj100@tcp(localhost:3306)/post_server"
	createStmt = `CREATE TABLE IF NOT EXISTS posts (
		postid CHAR(32) PRIMARY KEY,
		title VARCHAR(250),
		description TEXT,
		created_by CHAR(32),
		created_at BIGINT,
		updated_at BIGINT
	)`
)

type server struct {
	pb.UnimplementedPostServiceServer
	db *sql.DB
}

func main() {
	// Connect to MySQL
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to MySQL: %v", err)
	}
	defer db.Close()

	// Create table if not exists
	_, err = db.Exec(createStmt)
	if err != nil {
		log.Fatalf("failed to create table: %v", err)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, &server{db: db})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	// Validate input
	if req.Post.GetPostid() == "" || req.Post.GetTitle() == "" || len(req.Post.GetTitle()) > 100 || req.Post.GetDescription() == "" {
		return nil, fmt.Errorf("validation error")
	}

	// Insert into database
	_, err := s.db.Exec("INSERT INTO posts (postid, title, description, created_by, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		req.Post.GetPostid(),
		req.Post.GetTitle(),
		req.Post.GetDescription(),
		req.Post.GetCreatedBy(),
		req.Post.GetCreatedAt().Seconds,
		req.Post.GetUpdatedAt().Seconds,
	)
	if err != nil {
		return nil, err
	}

	return &pb.CreatePostResponse{Post: req.Post}, nil
}

func (s *server) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	row := s.db.QueryRow("SELECT postid, title, description, created_by, created_at, updated_at FROM posts WHERE postid = ?", req.GetPostid())
	post := &pb.Post{}
	var createdAt, updatedAt int64
	if err := row.Scan(&post.Postid, &post.Title, &post.Description, &post.CreatedBy, &createdAt, &updatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("post not found")
		}
		return nil, err
	}
	post.CreatedAt = timestamppb.New(time.Unix(createdAt, 0))
	post.UpdatedAt = timestamppb.New(time.Unix(updatedAt, 0))
	return &pb.GetPostResponse{Post: post}, nil
}

func (s *server) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
	// Validate input
	if req.Post.GetPostid() == "" || req.Post.GetTitle() == "" || len(req.Post.GetTitle()) > 100 || req.Post.GetDescription() == "" {
		return nil, fmt.Errorf("validation error")
	}

	// Update database
	_, err := s.db.Exec("UPDATE posts SET title = ?, description = ?, created_by = ?, created_at = ?, updated_at = ? WHERE postid = ?",
		req.Post.GetTitle(),
		req.Post.GetDescription(),
		req.Post.GetCreatedBy(),
		req.Post.GetCreatedAt().Seconds,
		req.Post.GetUpdatedAt().Seconds,
		req.Post.GetPostid(),
	)
	if err != nil {
		return nil, err
	}

	return &pb.UpdatePostResponse{Post: req.Post}, nil
}

func (s *server) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	_, err := s.db.Exec("DELETE FROM posts WHERE postid = ?", req.GetPostid())
	if err != nil {
		return nil, err
	}
	return &pb.DeletePostResponse{}, nil
}
