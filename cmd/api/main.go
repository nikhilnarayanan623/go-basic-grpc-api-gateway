package main

import (
	"fmt"
	"log"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/config"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("faild to load config error:%v", err.Error())
	}
	fmt.Println(cfg)
}
