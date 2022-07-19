package service

import "go-hospital-server/internal/core/repository"

type Service struct {
	Auth       *AuthService
	Patient    *PatientService
	OutPatient *OutPatientService
	User       *UserService
	Facility   *FacilityService
	Schedule   *ScheduleService
	Session    *SessionService
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Auth:       NewAuthService(r.Auth),
		Patient:    NewPatientService(r.Patient),
		OutPatient: NewOutPatientService(r.OutPatient),
		User:       NewUserService(r.User),
		Facility:   NewFacilityService(r.Facility),
		Schedule:   NewScheduleService(r.Schedule),
		Session:    NewSessionService(r.Session),
	}
}
