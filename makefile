proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/*.proto