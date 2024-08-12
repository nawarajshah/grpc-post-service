tidy:
	go mod tidy

gen:
	protoc -I. -Igoogleapis --go_out=. --go-grpc_out=. pb/post.proto
	protoc -I. -Igoogleapis --go_out=. --go-grpc_out=. pb/post_request.proto
	protoc -I. -Igoogleapis --go_out=. --go-grpc_out=. pb/post_response.proto
	protoc -I. -Igoogleapis --go_out=. --go-grpc_out=. pb/post_service.proto
	protoc -I. -Igoogleapis --go_out=. --go-grpc_out=. pb/comment.proto
	protoc -I. -Igoogleapis --go_out=. --go-grpc_out=. pb/comment_request.proto
	protoc -I. -Igoogleapis --go_out=. --go-grpc_out=. pb/comment_response.proto
	protoc -I. -Igoogleapis --go_out=. --go-grpc_out=. pb/comment_service.proto

clean:
	del .\pb\*.go

runAPI:
	go run post-api/main.go

runService:
	go run post-service/cmd/main.go
