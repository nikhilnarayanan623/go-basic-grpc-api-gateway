package clients

import (
	"fmt"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/clients/interfaces"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/config"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/domain"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type userServiceClient struct {
	client pb.UserServiceClient
}

func NewUserServiceClient(cfg *config.Config) (interfaces.UserServiceClient, error) {

	gc, err := grpc.Dial(cfg.UserServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewUserServiceClient(gc)

	return &userServiceClient{
		client: client,
	}, nil
}

func (c *userServiceClient) GetProfile(ctx context.Context, userID uint32) (domain.UserProfile, error) {

	res, err := c.client.GetUserProfile(ctx, &pb.GetUserProfileRequest{
		UserId: userID,
	})

	if err != nil {
		return domain.UserProfile{}, err
	}

	if res != nil && res.Response.Error != "" {
		return domain.UserProfile{}, fmt.Errorf("%s", res.Response.Error)
	}

	return domain.UserProfile{
		ID:        res.UserId,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Email:     res.Email,
	}, nil
}
