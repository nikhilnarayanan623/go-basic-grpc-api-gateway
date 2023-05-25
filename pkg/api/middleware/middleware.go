package middleware

import (
	"net/http"
	clientInterface "nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/clients/interfaces"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/domain"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middleware interface {
	UserAuthenticate(ctx *gin.Context)
}

type middleware struct {
	client clientInterface.AuthServiceClient
}

func NewMiddleware(client clientInterface.AuthServiceClient) Middleware {
	return &middleware{
		client: client,
	}
}

func (c *middleware) UserAuthenticate(ctx *gin.Context) {

	authorzation := ctx.GetHeader("authorization")

	authFields := strings.Fields(authorzation)

	if len(authFields) < 2 || authFields[1] == "" {
		response := utils.ErrorResponse(http.StatusBadRequest, "faild to get token", "no token found", nil)
		ctx.JSON(http.StatusBadGateway, response)
		return
	}

	accessToken := authFields[1]

	res, err := c.client.ValidateAccessToken(ctx, domain.ValidatTokenRequest{
		AccessToken: accessToken,
	})

	if err != nil {
		response := utils.ErrorResponse(http.StatusBadGateway, "faild to validate token", res.Error, nil)
		ctx.JSON(http.StatusBadGateway, response)
		return
	}

	ctx.Set("userId", res.UserID)
	ctx.Next()
}
