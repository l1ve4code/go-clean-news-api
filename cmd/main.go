package main

import (
	"go-clean-news-api/internal/adapters/routes"
	"go-clean-news-api/internal/composites"
	"go-clean-news-api/internal/config"
)

func main() {

	MySQL := composites.NewMySQLComposite(config.AppConfig.ConnectionString)
	UserComposite := composites.NewUserComposite(MySQL)
	ControllerComposite := composites.NewControllerComposite(UserComposite)

	routes.InitRoutes(ControllerComposite).Run(":8080")

}
