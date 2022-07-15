package routes

import (
	"go-hospital-server/internal/framework/transport/controller"
	mw "go-hospital-server/internal/framework/transport/middleware"

	"github.com/labstack/echo/v4"
)


func NewFacilityRoutes(e *echo.Group, acon *controller.FacilityController, middleware ...echo.MiddlewareFunc) {
	patient := e.Group("/facility", middleware...)
	patient.Use(mw.AdminPermission)
	patient.POST("", acon.Create)
	patient.GET("", acon.GetAllFacility)
	patient.GET("/:id", acon.GetFacilityByID)
	patient.PUT("/:id/update", acon.Update)
	patient.DELETE("/:id/delete", acon.Delete)
}
