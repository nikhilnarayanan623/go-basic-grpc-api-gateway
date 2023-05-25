package main

import (
	"log"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/config"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/di"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("faild to load config error:%s", err.Error())
	}

	server, err := di.InitializeApi(cfg)
	if err != nil {
		log.Fatalf("faild to initialize api error:%s", err.Error())
	}

	server.Start()
}
