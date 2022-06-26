package routes

import (
	"go-hospital-server/internal/framework/transport/controller"
	mw "go-hospital-server/internal/framework/transport/middleware"

	"github.com/labstack/echo/v4"
)

func NewOutPatientRoutes(e *echo.Group, acon *controller.OutPatientController, middleware ...echo.MiddlewareFunc) {
	patient := e.Group("/outpatient", middleware...)
	patient.GET("", acon.GetAllOutPatient)
	patient.GET("/:id", acon.FindByID)
	patient.POST("", acon.NewMedicRecord, mw.AdminPermission)
	patient.GET("/report", acon.Report, mw.AdminPermission)
	patient.GET("/log", acon.ReportLog, mw.DoctorNursePermission)
	patient.POST("/:id/process", acon.Process)
}
