package interfaces

import (
	"context"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/domain"
)

type UserServiceClient interface {
	GetProfile(ctx context.Context, userID uint32) (domain.UserProfile, error)
}
