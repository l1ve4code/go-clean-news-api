package user

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	RegisterUser(context *gin.Context)
	GenerateToken(context *gin.Context)
	Ping(context *gin.Context)
}