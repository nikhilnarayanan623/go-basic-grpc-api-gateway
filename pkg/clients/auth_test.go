package clients

// import (
// 	"context"
// 	"net"
// 	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/config"
// 	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/mocks"
// 	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/pb/authpb"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// 	"google.golang.org/grpc"
// )

// func TestUserSignup(t *testing.T) {

// 	testCases := []struct {
// 		name      string
// 		buildStub func(authServer *mocks.MockAuthServiceServer)
// 	}{
// 		{
// 			name:      "Test1",
// 			buildStub: func(authServer *mocks.MockAuthServiceServer) {},
// 		},
// 	}

// 	for _, test := range testCases {

// 		t.Run(test.name, func(t *testing.T) {
// 			ctl := gomock.NewController(t)
// 			authServer := mocks.NewMockAuthServiceServer(ctl)

// 			grpcServer := grpc.NewServer()
// 			authpb.RegisterAuthServiceServer(grpcServer, &Test{})
// 			test.buildStub(authServer)

// 			lis, err := net.Listen("tcp", ":50051")
// 			assert.NoError(t, err)
// 			defer lis.Close()

// 			assert.NoError(t, err)

// 			NewAuthServiceClient(&config.Config{
// 				AuthServiceUrl: "500",
// 			})
// 		})
// 	}

// }

// type Test struct {
// 	authpb.AuthServiceServer
// }

// func (c *Test) UserLogin(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
// 	return &authpb.LoginResponse{}, nil
// }

// func (c *Test) UserSignup(ctx context.Context, req *authpb.SignupRequest) (*authpb.SignupRsponse, error) {
// 	return &authpb.SignupRsponse{}, nil
// }

// func (c *Test) ValidateAccessToken(ctx context.Context, req *authpb.ValidateRequest) (*authpb.ValidateResponse, error) {
// 	return &authpb.ValidateResponse{}, nil
// }

// func TestUserLogin(t *testing.T) {

// }
