package controller

import "go-hospital-server/internal/core/service"

type Controller struct {
	Auth *AuthController
	Patient *PatientController
	OutPatient *OutPatientController
}

func NewController(srv *service.Service) *Controller {
	return &Controller{
		Auth: NewAuthController(srv.Auth),
		Patient: NewPatientController(srv.Patient),
		OutPatient: NewOutPatientController(srv.OutPatient),
	}
}
