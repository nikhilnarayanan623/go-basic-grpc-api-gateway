
run: 
	go run cmd/api/main.go

proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/*.proto

mock:
	mockgen -destination=./pkg/mocks/mock_service.go -package=mocks \
	nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/pb/authpb AuthServiceServer

wire:
	cd pkg/di && wire

all : proto wire run