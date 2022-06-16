package routes

import (
	"go-hospital-server/internal/framework/transport/controller"
	mw "go-hospital-server/internal/framework/transport/middleware"

	"github.com/labstack/echo/v4"
)

func NewPatientRoutes(e *echo.Group, acon *controller.PatientController, middleware ...echo.MiddlewareFunc) {
	patient := e.Group("/patient", middleware...)
	patient.POST("", acon.CreatePatient, mw.AdminPermission)
	patient.GET("", acon.GetAllPatient)
	patient.GET("/:id", acon.GetPatientByID)
	patient.PUT("/:id/update", acon.UpdatePatient, mw.AdminPermission)
	patient.DELETE("/:id/delete", acon.DeletePatient, mw.AdminPermission)
}
