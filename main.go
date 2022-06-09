package main

import (
	"go-hospital-server/internal/core/service"
	"go-hospital-server/internal/framework/database"
	"go-hospital-server/internal/framework/repository"
	"go-hospital-server/internal/framework/routes"
	"go-hospital-server/internal/framework/transport/controller"
	"go-hospital-server/internal/framework/transport/middleware"
	"go-hospital-server/internal/utils/config"
	"go-hospital-server/internal/utils/logger"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           Holy Hospital Sever API
// @version         1.0
// @description     server API for Holy Hospital Application.

// @securityDefinitions.apikey ApiKey
// @in header
// @name Authorization

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api
// @schemes http
func main() {

	config.LoadConfig()

	db, mongodb := database.InitDatabase()
	repo := repository.NewRepository(db, mongodb)
	serv := service.NewService(repo)
	ctrl := controller.NewController(serv)
	logger.NewLogger(mongodb)

	e := echo.New()
	e.GET("/*", echoSwagger.WrapHandler)

	api := e.Group("/api")
	middleware.NewJWTConnection(mongodb)
	routes.NewRoutes(api, ctrl, middleware.JWT)

	middleware.Logging(e)

	e.Start(":" + config.SERVER_PORT)
}
