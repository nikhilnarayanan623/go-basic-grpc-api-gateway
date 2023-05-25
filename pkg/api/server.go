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

func NewServer(cfg *config.Config, autHandler interfaces.AuthHandler, middleware middleware.Middleware) *Server {

	engine := gin.New()

	auth := engine.Group("/auth", middleware.UserAuthenticate)
	auth.POST("/signup", autHandler.UserSignup)
	auth.POST("/login", autHandler.UserLogin)

	return &Server{
		engine: engine,
		cfg:    cfg,
	}
}

func (s *Server) Start() {
	s.engine.Run(s.cfg.Port)
}
