package user

import "github.com/gin-gonic/gin"

type Service interface {
	RegisterUser(context *gin.Context, u *User) *User
	GenerateToken(password string, email string) string
}