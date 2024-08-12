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

	// Ensure the comments table exists
	ensureCommentsTableExists(db)

	return db
}

func ensureCommentsTableExists(db *sql.DB) {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS comments (
		commentid CHAR(32) PRIMARY KEY,
		postid CHAR(32) NOT NULL,
		userid CHAR(32) NOT NULL,
		content TEXT NOT NULL,
		created_at BIGINT NOT NULL,
		updated_at BIGINT NOT NULL,
		FOREIGN KEY (postid) REFERENCES posts(postid) ON DELETE CASCADE
	);`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Error creating comments table: %v", err)
	}

	fmt.Println("Comments table ensured to exist.")
}
