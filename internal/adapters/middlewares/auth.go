package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-clean-news-api/internal/adapters/auth"
	"strings"
)

func Auth() gin.HandlerFunc{
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		tokenString = strings.Split(tokenString, "Bearer ")[1]
		err:= auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}