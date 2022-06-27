package user

import (
	"github.com/gin-gonic/gin"
	"go-clean-news-api/internal/domain/entity"
)

type Service interface {
	RegisterUser(context *gin.Context, u *entity.User) *entity.User
	GenerateToken(password string, email string) string
}