tidy:
	go mod tidy

gen:
	protoc --go_out=. --go-grpc_out=. pb/post.proto

clean:
	del .\pb\*.go

runClient:
	go run client/main.go

runServer:
	go run server/main.go