package http

import (
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/api/handler/interfaces"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/api/middleware"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	cfg    *config.Config
}

func NewServer(cfg *config.Config, autHandler interfaces.AuthHandler, middleware middleware.Middleware,
	userHandler interfaces.UserHandler) *Server {

	engine := gin.Default()

	auth := engine.Group("/auth")
	auth.POST("/signup", autHandler.UserSignup)
	auth.POST("/login", autHandler.UserLogin)

	user := engine.Group("/user")

	user.Use(middleware.UserAuthenticate)

	user.GET("/profile", userHandler.GetProfile)

	return &Server{
		engine: engine,
		cfg:    cfg,
	}
}

func (s *Server) Start() {
	s.engine.Run(s.cfg.Port)
}
