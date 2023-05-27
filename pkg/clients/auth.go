package clients

import (
	"context"
	"fmt"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/clients/interfaces"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/config"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/domain"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authServiceClient struct {
	client pb.AuthServiceClient
}

func NewAuthServiceClient(cfg *config.Config) (interfaces.AuthServiceClient, error) {

	grpcClientConn, err := grpc.Dial(cfg.AuthServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc connection error:%s", err.Error())
	}

	client := pb.NewAuthServiceClient(grpcClientConn)

	return &authServiceClient{
		client: client,
	}, nil
}

func (c *authServiceClient) UserSignup(ctx context.Context, signupDetails domain.SignupRequest) error {

	res, err := c.client.UserSignup(context.Background(), &pb.SignupRequest{
		Email:     signupDetails.Email,
		Password:  signupDetails.Password,
		FirstName: signupDetails.FirstName,
		LastName:  signupDetails.LastName,
	})
	if err != nil {
		return err
	}
	if res.Response.Error != "" {
		return fmt.Errorf("%s", res.Response.Error)
	}
	return nil
}

func (c *authServiceClient) UserLogin(ctx context.Context, loginDetails domain.LoginRequest) (domain.LoginResponse, error) {

	res, err := c.client.UserLogin(context.Background(), &pb.LoginRequest{
		Email:    loginDetails.Email,
		Password: loginDetails.Password,
	})

	if err != nil {
		fmt.Println("first", err.Error())
		return domain.LoginResponse{}, err
	}

	if res.Response.Error != "" {
		return domain.LoginResponse{}, fmt.Errorf("%s", res.Response.Error)
	}

	return domain.LoginResponse{
		AccessToken: res.AccessToken,
	}, nil
}

func (c *authServiceClient) ValidateAccessToken(ctx context.Context, req domain.ValidateTokenRequest) (domain.ValidateTokenResponse, error) {

	res, err := c.client.ValidateAccessToken(context.Background(), &pb.ValidateRequest{
		AccessToken: req.AccessToken,
	})
	fmt.Println("res", res)
	if err != nil {
		return domain.ValidateTokenResponse{}, err
	}
	fmt.Println("res", res)

	if res.Response.Error != "" {
		return domain.ValidateTokenResponse{}, fmt.Errorf("%s", res.Response.Error)
	}
	fmt.Println("res", res.UserId)

	return domain.ValidateTokenResponse{
		UserID: res.UserId,
	}, nil
}
