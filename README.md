# gRPC Post and Comment Service

This project is a gRPC-based service with a RESTful API, using the Gin framework for managing posts and comments. It connects to a MySQL database and allows clients to perform CRUD operations on posts and comments.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Setup](#setup)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Testing with Thunder Client](#testing-with-thunder-client)
- [Environment Variables](#environment-variables)

## Features

- Create, read, update, and delete posts
- Create, read, update, and delete comments associated with posts
- RESTful API endpoints for interaction using Gin
- Automatic database table creation

## Tech Stack

- **Go**: Programming language
- **gRPC**: Communication protocol
- **Gin**: Web framework for building RESTful APIs
- **MySQL**: Database
- **Protobuf**: Serialization format
- **godotenv**: Environment variable management

## Setup

### Prerequisites

- Go 1.20 or later
- MySQL server
- Protobuf compiler (`protoc`) and Go plugins

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/nawarajshah/grpc-post-service.git
   cd grpc-post-service
   ```

2. **Install dependencies**:

   ```bash
   go mod tidy
   ```

3. **Install Protobuf tools**:

   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.53.0
   ```

   Make sure these binaries are in your `PATH`.

4. **Create a `.env` file**:

   Create a `.env` file in the root directory with the following content:

   ```env
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_HOST=localhost
   DB_PORT=3306
   DB_NAME=post_server
   ```

5. **Generate Protobuf files**:

   ```bash
   make gen
   ```

## Running the Application

### Start the gRPC Server

Navigate to the `post-service` directory and run:

```bash
make runService
```

This command will start the gRPC server and automatically create the required database tables if they don't exist.

### Start the REST API

Open a new terminal, navigate to the `post-api` directory, and run:

```bash
make runAPI
```

This command will start the REST API server using the Gin framework.

## API Endpoints

Here are the available endpoints for managing posts and comments:

### Posts

- **Create Post**: `POST /api/posts`
- **Get Post**: `GET /api/posts/:postId`
- **Update Post**: `PUT /api/posts/:postId`
- **Delete Post**: `DELETE /api/posts/:postId`

### Comments

- **Create Comment**: `POST /api/posts/:postId/comments`
- **Get Comment**: `GET /api/posts/:postId/comments/:commentId`
- **Update Comment**: `PUT /api/posts/:postId/comments/:commentId`
- **Delete Comment**: `DELETE /api/posts/:postId/comments/:commentId`
- **List Comments**: `GET /api/posts/:postId/comments`

## Testing with Thunder Client

1. **Create a Comment**:

   - **Method**: POST
   - **URL**: `http://localhost:8080/api/posts/{postId}/comments`
   - **Body**:

     ```json
     {
       "comment": {
         "commentId": "12345",
         "userId": "user123",
         "content": "This is a test comment."
       }
     }
     ```

2. **Get a Comment**:

   - **Method**: GET
   - **URL**: `http://localhost:8080/api/posts/{postId}/comments/{commentId}`

3. **Update a Comment**:

   - **Method**: PUT
   - **URL**: `http://localhost:8080/api/posts/{postId}/comments/{commentId}`
   - **Body**:

     ```json
     {
       "comment": {
         "commentId": "12345",
         "userId": "user123",
         "content": "This is an updated test comment."
       }
     }
     ```

4. **Delete a Comment**:

   - **Method**: DELETE
   - **URL**: `http://localhost:8080/api/posts/{postId}/comments/{commentId}`
   - **Headers**:
     - **userId**: `user123`

5. **List Comments**:

   - **Method**: GET
   - **URL**: `http://localhost:8080/api/posts/{postId}/comments`

## Environment Variables

- **`DB_USER`**: Database user name
- **`DB_PASSWORD`**: Database password
- **`DB_HOST`**: Database host (e.g., `localhost`)
- **`DB_PORT`**: Database port (default is `3306` for MySQL)
- **`DB_NAME`**: Database name (e.g., `post_server`)

### Important Notes

- **Database**: Ensure that the MySQL server is running and accessible with the credentials specified in the `.env` file.
- **Security**: This project is a basic setup for educational purposes. For production, ensure to implement proper authentication, authorization, and security measures.
