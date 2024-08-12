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

<<<<<<<< HEAD:post-service/cmd/main.go
	// Initialize the repository
========
	// Create the PostRepository instance
>>>>>>>> 64d27f067b21c91a8acc82b612efb223375b5709:post-service/main.go
	postRepo := repo.NewPostRepository(database)

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

<<<<<<<< HEAD:post-service/cmd/main.go
	// Register the PostServiceServer
========
	// Register the PostServiceServer using the repository
>>>>>>>> 64d27f067b21c91a8acc82b612efb223375b5709:post-service/main.go
	postService := service.NewPostServiceServer(postRepo)
	pb.RegisterPostServiceServer(grpcServer, postService)

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
