package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/nawarajshah/grpc-post-service/postpb"
	"github.com/nawarajshah/grpc-post-service/server"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	postpb.RegisterPostServiceServer(s, &server.Server{})

	log.Println("Server is running on port :50051")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
