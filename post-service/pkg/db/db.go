package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() *sql.DB {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get database credentials from environment variables
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	// Create the Data Source Name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)

	// Connect to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	fmt.Println("Connected to the database successfully")

	// Ensure all required tables exist
	ensureTablesExist(db)

	return db
}

func ensureTablesExist(db *sql.DB) {
	// Ensure the users table exists
	createUsersTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		user_id CHAR(36) PRIMARY KEY,
		email VARCHAR(191) UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		is_email_verified BOOLEAN NOT NULL DEFAULT FALSE,
		verification_code CHAR(6),
		created_at BIGINT NOT NULL
	);
	`
	_, err := db.Exec(createUsersTableQuery)
	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}

	// Ensure the posts table exists
	createPostsTableQuery := `
	CREATE TABLE IF NOT EXISTS posts (
		postid CHAR(32) PRIMARY KEY,
		title VARCHAR(250) NOT NULL,
		description TEXT,
		created_by CHAR(36) NOT NULL,
		created_at BIGINT NOT NULL,
		updated_at BIGINT NOT NULL,
		FOREIGN KEY (created_by) REFERENCES users(user_id) ON DELETE CASCADE
	);
	`
	_, err = db.Exec(createPostsTableQuery)
	if err != nil {
		log.Fatalf("Error creating posts table: %v", err)
	}

	// Ensure the comments table exists
	createCommentsTableQuery := `
	CREATE TABLE IF NOT EXISTS comments (
		commentid CHAR(32) PRIMARY KEY,
		postid CHAR(32) NOT NULL,
		userid CHAR(36) NOT NULL,
		content TEXT NOT NULL,
		created_at BIGINT NOT NULL,
		updated_at BIGINT NOT NULL,
		FOREIGN KEY (postid) REFERENCES posts(postid) ON DELETE CASCADE,
		FOREIGN KEY (userid) REFERENCES users(user_id) ON DELETE CASCADE
	);
	`
	_, err = db.Exec(createCommentsTableQuery)
	if err != nil {
		log.Fatalf("Error creating comments table: %v", err)
	}

	fmt.Println("All required tables ensured to exist.")
}
