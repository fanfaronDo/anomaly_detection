build_client:
	@go build -o client ./cmd/client/main.go

build_server:
	@go build -o server ./cmd/server/main.go

build_test:
	@go build -o test ./cmd/test/main.go

proto_generate:
	@protoc --go_out=pkg/api --go_opt=paths=source_relative --go-grpc_out=pkg/api --go-grpc_opt=paths=source_relative ./api/proto/serv.proto


clean:
	rm ./server ./client