package routes

import (
	"go-hospital-server/internal/framework/transport/controller"
	mw "go-hospital-server/internal/framework/transport/middleware"

	"github.com/labstack/echo/v4"
)


func NewUserRoutes(e *echo.Group, acon *controller.UserController, middleware ...echo.MiddlewareFunc) {
	patient := e.Group("/user", middleware...)
	patient.Use(mw.AdminPermission)
	patient.POST("", acon.Create)
	patient.GET("", acon.GetAllUser)
	patient.GET("/:id", acon.GetUserByID)
	patient.PUT("/:id/update", acon.Update)
	patient.DELETE("/:id/delete", acon.Delete, mw.AdminPermission)
}
