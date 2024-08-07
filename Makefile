tidy:
	go mod tidy

gen:
	protoc --go_out=pb --go_opt=paths=source_relative \
       --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
       --proto_path=proto \
       proto/post.proto

clean:
	del .\pb\*.go

run:
	go run main.go