package controller

import "go-hospital-server/internal/core/service"

type Controller struct {
	Auth       *AuthController
	Patient    *PatientController
	OutPatient *OutPatientController
	User       *UserController
	Facility   *FacilityController
	Schedule   *ScheduleController
	Session    *SessionController
}

func NewController(srv *service.Service) *Controller {
	return &Controller{
		Auth:       NewAuthController(srv.Auth),
		Patient:    NewPatientController(srv.Patient),
		OutPatient: NewOutPatientController(srv.OutPatient),
		User:       NewUserController(srv.User),
		Facility:   NewFacilityController(srv.Facility),
		Schedule:   NewScheduleController(srv.Schedule),
		Session:    NewSessionController(srv.Session),
	}
}
