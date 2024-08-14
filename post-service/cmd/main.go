package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/db"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/repo"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/service"
)

func main() {
	// Connect to the database
	database := db.Connect()
	defer database.Close()

	// Initialize repositories
	postRepo := repo.NewPostRepository(database)
	commentRepo := repo.NewCommentRepository(database)
	userRepo := repo.NewUserRepository(database)

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the PostServiceServer
	postService := service.NewPostServiceServer(postRepo)
	pb.RegisterPostServiceServer(grpcServer, postService)

	// Register the CommentServiceServer
	commentService := service.NewCommentServiceServer(commentRepo)
	pb.RegisterCommentServiceServer(grpcServer, commentService)

	// Register the AuthServiceServer
	authService := service.NewAuthServiceServer(userRepo)
	pb.RegisterAuthServiceServer(grpcServer, authService) // Ensure this line is present

	// Listen on port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	fmt.Println("Server is listening on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
