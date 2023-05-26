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

func (c *authServiceClient) UserSignup(ctx context.Context, signupDetails domain.SignupRequest) (domain.SignupRsponse, error) {

	response, err := c.client.UserSignup(context.Background(), &pb.SignupRequest{
		Email:     signupDetails.Email,
		Password:  signupDetails.Password,
		FirstName: signupDetails.FristName,
		LastName:  signupDetails.LastName,
	})
	if err != nil {
		return domain.SignupRsponse{}, err
	}
	return domain.SignupRsponse{
		StatusCode: response.Response.StatusCode,
		Error:      response.Response.Error,
		Message:    response.Response.Message,
	}, nil
}

func (c *authServiceClient) UserLogin(ctx context.Context, loginDetails domain.LoginRequest) (domain.LoginResponse, error) {

	response, err := c.client.UserLogin(context.Background(), &pb.LoginRequest{
		Email:    loginDetails.Email,
		Password: loginDetails.Password,
	})

	if err != nil {
		return domain.LoginResponse{}, err
	}

	return domain.LoginResponse{
		StatusCode:  response.Response.StatusCode,
		Message:     response.Response.Message,
		Error:       "",
		AccessToken: response.AccessToken,
	}, nil
}

func (c *authServiceClient) ValidateAccessToken(ctx context.Context, req domain.ValidatTokenRequest) (domain.ValidateTokenResponse, error) {

	res, err := c.client.ValidateAccessToken(context.Background(), &pb.ValidateRequest{
		AccessToken: req.AccessToken,
	})

	if err != nil {
		return domain.ValidateTokenResponse{}, err
	}

	return domain.ValidateTokenResponse{
		StatusCode: res.Response.StatusCode,
		Message:    res.Response.Message,
		Error:      "",
		UserID:     res.UserId,
	}, nil
}
