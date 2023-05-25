package clients

import (
	"context"
	"fmt"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/clients/interfaces"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/config"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/domain"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/pb/authpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authServiceClient struct {
	client authpb.AuthServiceClient
}

func NewAuthServiceClient(cfg *config.Config) (interfaces.AuthServiceClient, error) {

	grpcClientConn, err := grpc.Dial(cfg.AuthServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("faild to create grpc connection error:%s", err.Error())
	}

	authSericeClient := authpb.NewAuthServiceClient(grpcClientConn)

	return &authServiceClient{
		client: authSericeClient,
	}, nil
}

func (c *authServiceClient) UserSignup(ctx context.Context, signupDetails domain.SignupRequest) (domain.SignupRsponse, error) {

	response, err := c.client.UserSignup(context.Background(), &authpb.SignupRequest{
		Email:     signupDetails.Email,
		Password:  signupDetails.Password,
		FirstName: signupDetails.FristName,
		LastName:  signupDetails.LastName,
	})

	if err != nil {
		return domain.SignupRsponse{}, err
	}
	return domain.SignupRsponse{
		StatusCode: response.Status.StatusCode,
		Error:      "",
		Message:    response.Status.Message,
	}, nil
}

func (c *authServiceClient) UserLogin(ctx context.Context, loginDetails domain.LoginRequest) (domain.LoginResponse, error) {

	response, err := c.client.UserLogin(context.Background(), &authpb.LoginRequest{
		Email:    loginDetails.Email,
		Password: loginDetails.Password,
	})

	if err != nil {
		return domain.LoginResponse{}, err
	}

	return domain.LoginResponse{
		StatusCode:  response.Status.StatusCode,
		Message:     response.Status.Message,
		Error:       "",
		AccessToken: response.AccessToken,
	}, nil
}

func (c *authServiceClient) ValidateAccessToken(ctx context.Context, req domain.ValidatTokenRequest) (domain.ValidateTokenResponse, error) {

	res, err := c.client.ValidateAccessToken(context.Background(), &authpb.ValidateRequest{
		AccessToken: req.AccessToken,
	})

	if err != nil {
		return domain.ValidateTokenResponse{}, err
	}

	return domain.ValidateTokenResponse{
		StatusCode: res.Status.StatusCode,
		Message:    res.Status.Message,
		Error:      "",
		UserID:     res.UserId,
	}, nil
}
