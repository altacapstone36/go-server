package main

import (
	"go-hospital-server/internal/core/service"
	"go-hospital-server/internal/framework/database"
	"go-hospital-server/internal/framework/repository"
	"go-hospital-server/internal/framework/routes"
	"go-hospital-server/internal/framework/transport/controller"
	mw "go-hospital-server/internal/framework/transport/middleware"
	"go-hospital-server/internal/utils/config"
	"go-hospital-server/internal/utils/logger"
	"go-hospital-server/internal/utils/validators"

	_ "go-hospital-server/docs"

	"github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
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

// @host     go-hospital-server.herokuapp.com
// @BasePath  /api
// @schemes http https
func main() {

	config.LoadConfig()

	db, mongodb := database.InitDatabase()
	repo := repository.NewRepository(db, mongodb)
	serv := service.NewService(repo)
	ctrl := controller.NewController(serv)
	logger.NewLogger(mongodb)
	validators.NewValidator(db)

	e := echo.New()
	e.Use(emw.CORS())
	e.GET("/*", echoSwagger.WrapHandler)

	api := e.Group("/api")
	mw.NewJWTConnection(mongodb)
	routes.NewRoutes(api, ctrl, mw.JWT)

	mw.Logging(e)

	e.Start(":" + config.SERVER_PORT)
}
