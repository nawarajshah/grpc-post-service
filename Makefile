PROTO_DIR=pb
GAPI_DIR=googleapis

tidy:
	go mod tidy

gen:
	protoc -I. -I$(GAPI_DIR) --go_out=. --go-grpc_out=. $(PROTO_DIR)/post.proto
	protoc -I. -I$(GAPI_DIR) --go_out=. --go-grpc_out=. $(PROTO_DIR)/post_request.proto
	protoc -I. -I$(GAPI_DIR) --go_out=. --go-grpc_out=. $(PROTO_DIR)/post_response.proto
	protoc -I. -I$(GAPI_DIR) --go_out=. --go-grpc_out=. $(PROTO_DIR)/post_service.proto

clean:
	del .\pb\*.go

runAPI:
	go run post-api/main.go

runService:
	go run post-service/cmd/main.go
