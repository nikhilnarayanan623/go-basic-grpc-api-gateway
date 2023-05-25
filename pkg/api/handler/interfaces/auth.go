package interfaces

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	UserLogin(ctx *gin.Context)
	UserSignup(ctx *gin.Context)
}
