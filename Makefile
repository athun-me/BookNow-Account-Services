proto:
	protoc --go_out=. --go-grpc_out=. pkg/pb/user.proto

wire:
	go run github.com/google/wire/cmd/wire

run:
	go run cmd/api/main.go