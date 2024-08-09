tidy:
	go mod tidy

gen:
	protoc --go_out=. --go-grpc_out=. pb/post.proto
	protoc --go_out=. --go-grpc_out=. pb/post_request.proto
	protoc --go_out=. --go-grpc_out=. pb/post_response.proto
	protoc --go_out=. --go-grpc_out=. pb/post_service.proto

clean:
	del .\pb\*.go

runClient:
	go run client/main.go

runServer:
	go run server/cmd/main.go
