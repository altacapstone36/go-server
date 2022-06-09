package repository

import "go-hospital-server/internal/core/entity/models"

type PatientRepository interface {
	GetPatientByID(uint) (models.Patient, error)
	GetPatientByName(string) ([]models.Patient, error)
	GetAllPatient() ([]models.Patient, error)
	CreatePatient(models.Patient) (error)
	UpdatePatient(models.Patient) (error)
	DeletePatientByID(uint) (error)
}
