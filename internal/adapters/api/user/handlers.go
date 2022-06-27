package user

import (
	"github.com/gin-gonic/gin"
	"go-clean-news-api/internal/domain/user"
	"net/http"
)

type handler struct {
	userService user.Service
}

func NewHandler(service user.Service) Handler{
	return &handler{userService: service}
}

func (h *handler) RegisterUser(context *gin.Context){
	var user user.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	user = *h.userService.RegisterUser(context, &user)
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}

func (h *handler) GenerateToken(context *gin.Context){
	var request user.TokenRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	token := h.userService.GenerateToken(request.Password, request.Email)
	context.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *handler) Ping(context *gin.Context){
	context.JSON(http.StatusOK, gin.H{"token": "kek"})
}