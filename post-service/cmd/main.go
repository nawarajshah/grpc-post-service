package main

import (
	"log"
	"net"

	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/db"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/endpoint"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/facade"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/repo"
	"github.com/nawarajshah/grpc-post-service/post-service/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Connect to the database
	database := db.Connect()
	defer database.Close()

	// Initialize repositories
	userRepo := repo.NewUserRepository(database)
	verificationRepo := repo.NewVerificationRepository(database)
	postRepo := repo.NewPostRepository(database)
	commentRepo := repo.NewCommentRepository(database)

	// Initialize facades
	postFacade := facade.NewPostFacade(postRepo)

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Initialize service servers
	authService := service.NewAuthServiceServer(userRepo, verificationRepo)
	postService := endpoint.NewPostServiceRpc(postFacade) // Use the new endpoint
	commentService := service.NewCommentServiceServer(commentRepo, postRepo)

	// Register the services with the gRPC server
	pb.RegisterAuthServiceServer(grpcServer, authService)
	pb.RegisterPostServiceServer(grpcServer, postService)
	pb.RegisterCommentServiceServer(grpcServer, commentService)
	pb.RegisterVerificationServiceServer(grpcServer, service.NewVerificationServiceServer(verificationRepo))

	// Enable server reflection
	reflection.Register(grpcServer)

	// Listen on a TCP port
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
