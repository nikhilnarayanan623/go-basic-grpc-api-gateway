package middleware

import (
	"fmt"
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

	authorization := ctx.GetHeader("authorization")

	authFields := strings.Fields(authorization)
	fmt.Println("token", authFields)

	if len(authFields) < 2 {
		response := utils.ErrorResponse(http.StatusBadRequest, "failed verify request", "no token found", nil)
		ctx.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}

	accessToken := authFields[1]

	res, err := c.client.ValidateAccessToken(ctx, domain.ValidateTokenRequest{
		AccessToken: accessToken,
	})
	if err != nil {
		response := utils.ErrorResponse(http.StatusBadGateway, "failed verify request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}

	ctx.Set("userId", uint64(res.UserID))
	ctx.Next()
}
