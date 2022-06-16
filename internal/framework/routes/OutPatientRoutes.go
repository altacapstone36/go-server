package routes

import (
	"go-hospital-server/internal/framework/transport/controller"
	mw "go-hospital-server/internal/framework/transport/middleware"

	"github.com/labstack/echo/v4"
)

func NewOutPatientRoutes(e *echo.Group, acon *controller.OutPatientController, middleware ...echo.MiddlewareFunc) {
	patient := e.Group("/outpatient", middleware...)
	patient.GET("", acon.GetAllOutPatient)
	patient.POST("", acon.NewMedicRecord, mw.AdminPermission)
	patient.POST("/doctor", acon.NewMedicRecord, mw.DoctorPermission)
	patient.POST("/nurse", acon.NewMedicRecord, mw.NursePermission)
}
