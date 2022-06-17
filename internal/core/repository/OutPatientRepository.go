package repository

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"
)

type OutPatientRepository interface {
	NewMedicalRecord(models.MedicRecord) (error)
	Proceed(models.MedicRecord) (error)
	ListAvailable(int, string) ([]response.OutPatientResponse, error)
	FilterByDate(string, string) ([]response.OutPatientResponse, error)
}
