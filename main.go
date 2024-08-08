package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nawarajshah/grpc-post-service/postpb"
	"github.com/nawarajshah/grpc-post-service/server"
	"google.golang.org/grpc"
)

func main() {
	// Connect to MySQL
	dsn := "nawaraj:nawaraj100@tcp(localhost:3306)/post_server"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	defer db.Close()

	// Ensure the table exists
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS posts (
		postid CHAR(32) PRIMARY KEY,
		title VARCHAR(250),
		description TEXT,
		created_by CHAR(32),
		created_at BIGINT,
		updated_at BIGINT
	);`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	postpb.RegisterPostServiceServer(s, server.NewServer(db))

	log.Println("Server is running on port :50051")
	if err := s.Serve(listener); err != nil {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}
}