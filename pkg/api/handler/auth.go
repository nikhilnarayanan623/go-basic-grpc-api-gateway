package handler

import (
	"net/http"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/api/handler/interfaces"
	clientInterface "nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/clients/interfaces"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/domain"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authClient clientInterface.AuthServiceClient
}

func NewAuthHandler(client clientInterface.AuthServiceClient) interfaces.AuthHandler {
	return &authHandler{
		authClient: client,
	}
}

func (c *authHandler) UserLogin(ctx *gin.Context) {

	var body domain.LoginRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := utils.ErrorResponse(http.StatusBadRequest, "faild to bind inputs", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	res, err := c.authClient.UserLogin(ctx, body)

	if err != nil {
		errString := res.Error + err.Error()
		response := utils.ErrorResponse(http.StatusBadRequest, "faild to login", errString, nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.SuccessResponse(http.StatusOK, "successfully logged in", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *authHandler) UserSignup(ctx *gin.Context) {

	var body domain.SignupRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := utils.ErrorResponse(http.StatusBadRequest, "faild to bind inputs", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	res, err := c.authClient.UserSignup(ctx, body)
	if err != nil {
		errString := res.Error + err.Error()
		response := utils.ErrorResponse(http.StatusBadGateway, "faild to signup", errString, nil)
		ctx.JSON(http.StatusBadGateway, response)
		return
	}

	response := utils.SuccessResponse(http.StatusOK, "successfully signup completed", res)
	ctx.JSON(http.StatusOK, response)
}
