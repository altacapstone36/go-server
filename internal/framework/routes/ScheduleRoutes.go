package routes

import (
	"go-hospital-server/internal/framework/transport/controller"
	mw "go-hospital-server/internal/framework/transport/middleware"

	"github.com/labstack/echo/v4"
)

func NewScheduleRoutes(e *echo.Group, acon *controller.ScheduleController, middleware ...echo.MiddlewareFunc) {
	schedule := e.Group("/session/schedule", middleware...)
	schedule.Use(mw.AdminPermission)
	schedule.POST("", acon.Create)
	// schedule.PUT("/:sessionID/update", acon.Update)
	// schedule.DELETE("/:id/delete", acon.Delete, mw.AdminPermission)
}
