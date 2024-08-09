package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/nawarajshah/grpc-post-service/pb"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewPostServiceClient(conn)

	// Test CreatePost
	postId := "1234"
	createdAt := timestamppb.Now()
	post := &pb.Post{
		PostId:     postId,
		Title:      "My First Post",
		Description: "This is the content of my first post",
		CreatedBy:  "user123",
		CreatedAt:  createdAt,
		UpdatedAt:  createdAt,
	}

	createRes, err := client.CreatePost(context.Background(), &pb.CreatePostRequest{Post: post})
	if err != nil {
		log.Fatalf("Error creating post: %v", err)
	}
	fmt.Printf("Post created: %v\n", createRes.Post)

	// Test GetPost
	getRes, err := client.GetPost(context.Background(), &pb.GetPostRequest{PostId: postId})
	if err != nil {
		log.Fatalf("Error getting post: %v", err)
	}
	fmt.Printf("Post retrieved: %v\n", getRes.Post)

	// Test UpdatePost
	post.Title = "My Updated Post"
	post.UpdatedAt = timestamppb.Now()
	updateRes, err := client.UpdatePost(context.Background(), &pb.UpdatePostRequest{Post: post})
	if err != nil {
		log.Fatalf("Error updating post: %v", err)
	}
	fmt.Printf("Post updated: %v\n", updateRes.Post)

	// Test DeletePost
	_, err = client.DeletePost(context.Background(), &pb.DeletePostRequest{PostId: postId})
	if err != nil {
		log.Fatalf("Error deleting post: %v", err)
	}
	fmt.Println("Post deleted")
}
