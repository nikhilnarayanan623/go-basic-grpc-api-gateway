package handler

import (
	"net/http"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/api/handler/interfaces"
	clientInterface "nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/clients/interfaces"
	"nikhilnarayanan623/go-basic-grpc-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	client clientInterface.UserServiceClient
}

func NewUserHandler(client clientInterface.UserServiceClient) interfaces.UserHandler {
	return &userHandler{
		client: client,
	}
}

func (c *userHandler) GetProfile(ctx *gin.Context) {

	userID := ctx.GetUint64("userId")

	user, err := c.client.GetProfile(ctx, uint32(userID))

	if err != nil {
		response := utils.ErrorResponse(http.StatusInternalServerError, "failed get user details", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := utils.SuccessResponse(http.StatusOK, "successfully got user details", user)
	ctx.JSON(http.StatusOK, response)
}
