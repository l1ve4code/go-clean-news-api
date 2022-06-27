package routes

import (
	"github.com/gin-gonic/gin"
	"go-clean-news-api/internal/adapters/middlewares"
	"go-clean-news-api/internal/composites"
)

func InitRoutes(controllers *composites.ControllersComposite) *gin.Engine{
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.UserController.GenerateToken)
		api.POST("/user/register", controllers.UserController.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.UserController.Ping)
		}
	}
	return router
}