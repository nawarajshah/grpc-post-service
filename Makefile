# Go commands
GOCMD=go
GOTIDY=$(GOCMD) mod tidy
GORUN=$(GOCMD) run

# Protobuf generation commands
PROTOC=protoc
PROTO_FLAGS=--go_out=. --go-grpc_out=.
PROTO_DIR=pb

# Source directories
POST_API_DIR=post-api
POST_SERVICE_DIR=post-service

# List of proto files
PROTO_FILES=$(PROTO_DIR)/post.proto \
	$(PROTO_DIR)/post_request.proto \
	$(PROTO_DIR)/post_response.proto \
	$(PROTO_DIR)/post_service.proto \
	$(PROTO_DIR)/comment.proto \
	$(PROTO_DIR)/comment_request.proto \
	$(PROTO_DIR)/comment_response.proto \
	$(PROTO_DIR)/comment_service.proto \
	$(PROTO_DIR)/auth.proto \
	$(PROTO_DIR)/auth_request.proto \
	$(PROTO_DIR)/auth_response.proto \
	$(PROTO_DIR)/auth_service.proto \
	$(PROTO_DIR)/verification.proto \
	$(PROTO_DIR)/verification_request.proto \
	$(PROTO_DIR)/verification_response.proto \
	$(PROTO_DIR)/verification_service.proto

# Tidy up Go dependencies
tidy:
	$(GOTIDY)

# Generate protobuf files
gen:
	$(PROTOC) $(PROTO_FLAGS) $(PROTO_FILES)

# Clean generated files
clean:
	del .\pb\*.go

# Run the post-service gRPC server
runService:
	$(GORUN) $(POST_SERVICE_DIR)/cmd/main.go

# Run the post-api REST server
runAPI:
	$(GORUN) $(POST_API_DIR)/main.go
