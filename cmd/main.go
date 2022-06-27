package main

import (
	"go-clean-news-api/internal/adapters/routes"
	"go-clean-news-api/internal/composites"
	"go-clean-news-api/internal/config"
)

func main() {

	config.LoadAppConfig()
	MySQL := composites.NewMySQLComposite(config.AppConfig.ConnectionString)
	UserComposite := composites.NewUserComposite(MySQL)
	ControllerComposite := composites.NewControllerComposite(UserComposite)

	routes.InitRoutes(ControllerComposite).Run(config.AppConfig.Port)

}
