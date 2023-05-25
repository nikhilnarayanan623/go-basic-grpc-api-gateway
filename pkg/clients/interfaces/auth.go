package interfaces

import (
	"context"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/domain"
)

type AuthServiceClient interface {
	UserSignup(ctx context.Context, signupDeails domain.SignupRequest) (domain.SignupRsponse, error)
	UserLogin(ctx context.Context, loginDetails domain.LoginRequest) (domain.LoginResponse, error)
	ValidateAccessToken(ctx context.Context, tokenRequest domain.ValidatTokenRequest) (domain.ValidateTokenResponse, error)
}
