package main

import (
	"go-hospital-server/internal/core/service"
	"go-hospital-server/internal/framework/database"
	"go-hospital-server/internal/framework/repository"
	"go-hospital-server/internal/framework/routes"
	"go-hospital-server/internal/framework/transport/controller"
	"go-hospital-server/internal/framework/transport/middleware"
	"go-hospital-server/internal/utils"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           Question Board
// @version         1.0
// @description     server API for Question Board Application.

// @securityDefinitions.apikey ApiKey
// @in header
// @name Authorization

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api
// @schemes http
func main() {

	utils.LoadConfig()

	db := database.InitDatabase()
	repo := repository.NewRepository(db)
	serv := service.NewService(repo)
	ctrl := controller.NewController(serv)

	e := echo.New()
	e.GET("/*", echoSwagger.WrapHandler)

	api := e.Group("/api")
	routes.NewRoutes(api, ctrl, middleware.JWT())

	middleware.Logging(e)

	e.Start(":" + utils.SERVER_PORT)
}
