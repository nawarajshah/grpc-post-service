package main

import (
	"log"

	"github.com/nawarajshah/grpc-post-service/pb"
	"github.com/nawarajshah/grpc-post-service/post-api/controller"
	"github.com/nawarajshah/grpc-post-service/post-api/router"
	"github.com/nawarajshah/grpc-post-service/post-api/service"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Set up clients for the gRPC services
	postClient := pb.NewPostServiceClient(conn)
	commentClient := pb.NewCommentServiceClient(conn)

	// Create service instances
	postService := service.NewPostService(postClient)
	commentService := service.NewCommentService(commentClient)

	// Create controller instances
	postController := controller.NewPostController(postService)
	commentController := controller.NewCommentController(commentService)

	// Set up the Gin router
	r := router.SetupRouter(postController, commentController)

	// Run the Gin server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
