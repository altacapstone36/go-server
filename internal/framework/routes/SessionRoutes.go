package routes

import (
	"go-hospital-server/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewSessionRoutes(e *echo.Group, acon *controller.SessionController, middleware ...echo.MiddlewareFunc) {
	session := e.Group("/session", middleware...)
	session.GET("", acon.FindAllSession)
}
