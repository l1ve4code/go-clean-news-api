package user

import (
	"github.com/gin-gonic/gin"
	"go-clean-news-api/internal/adapters/auth"
	"go-clean-news-api/internal/adapters/db/user"
	"go-clean-news-api/internal/domain/entity"
	"net/http"
)

type service struct {
	storage user.Storage
}

func NewService(storage user.Storage) Service{
	return &service{storage: storage}
}

func (s *service) RegisterUser(context *gin.Context, user *entity.User) *entity.User{
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return nil
	}
	record := s.storage.Create(user)

	return record
}

func (s *service) GenerateToken(password string, email string) string{
	user := s.storage.FindByEmail(email)
	credentialError := user.CheckPassword(password)
	if credentialError != nil {
		return ""
	}
	tokenString, _:= auth.GenerateJWT(user.Email, user.Username)
	return tokenString
}