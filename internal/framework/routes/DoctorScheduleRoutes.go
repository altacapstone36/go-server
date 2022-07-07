package routes

import (
	"go-hospital-server/internal/framework/transport/controller"
	mw "go-hospital-server/internal/framework/transport/middleware"

	"github.com/labstack/echo/v4"
)

func NewDoctorScheduleRoutes(e *echo.Group, acon *controller.DoctorScheduleController, middleware ...echo.MiddlewareFunc) {
	schedule := e.Group("/schedule", middleware...)
	schedule.Use(mw.AdminPermission)
	schedule.GET("", acon.GetAllSchedule)
	schedule.POST("", acon.Create)
	schedule.PUT("/:sessionID/update", acon.Update)
	schedule.DELETE("/:id/delete", acon.Delete, mw.AdminPermission)
}
