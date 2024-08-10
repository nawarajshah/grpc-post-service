tidy:
	go mod tidy

gen:
	protoc --go_out=. --go-grpc_out=. pb/post.proto
	protoc --go_out=. --go-grpc_out=. pb/post_request.proto
	protoc --go_out=. --go-grpc_out=. pb/post_response.proto
	protoc --go_out=. --go-grpc_out=. pb/post_service.proto

clean:
	del .\pb\*.go

runAPI:
	go run post-api/main.go

runService:
	go run post-service/cmd/main.go
