//go:build wireinject
// +build wireinject

package di

import (
	http "nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/api"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/api/handler"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/api/middleware"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/clients"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/config"

	"github.com/google/wire"
)

func InitializeApi(cfg *config.Config) (*http.Server, error) {

	wire.Build(
		clients.NewAuthServiceClient,
		handler.NewAuthHandler,
		middleware.NewMiddleware,
		http.NewServer,
	)
	return &http.Server{}, nil
}
