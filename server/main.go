package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/server/pkg/db"
	"github.com/nawarajshah/grpc-post-service/server/pkg/service"
)

func main() {
	// Connect to the database
	database := db.Connect()
	defer database.Close()

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the PostServiceServer
	postService := service.NewPostServiceServer(database)
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
