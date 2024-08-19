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

	postClient := pb.NewPostServiceClient(conn)
	postService := service.NewPostService(postClient)
	postController := controller.NewPostController(postService)

	commentClient := pb.NewCommentServiceClient(conn)
	commentService := service.NewCommentService(commentClient)
	commentController := controller.NewCommentController(commentService)

	authClient := pb.NewAuthServiceClient(conn)
	authService := service.NewAuthService(authClient)
	authController := controller.NewAuthController(authService)

	verificationClient := pb.NewVerificationServiceClient(conn)
	verificationService := service.NewVerificationService(verificationClient)
	verificationController := controller.NewVerificationController(verificationService)

	// Set up the Gin router
	r := router.SetupRouter(postController, commentController, authController, verificationController)

	// Run the Gin server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
