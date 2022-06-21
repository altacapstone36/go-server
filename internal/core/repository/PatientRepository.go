package repository

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"
)

type PatientRepository interface {
	GetPatientByID(uint) (response.PatientDetails, error)
	GetPatientByName(string) ([]response.Patient, error)
	GetAllPatient() ([]response.Patient, error)
	CreatePatient(models.Patient) (error)
	UpdatePatient(models.Patient) (error)
	DeletePatientByID(uint) (error)
}
