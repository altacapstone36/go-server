package repository

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"
)

type OutPatientRepository interface {
	NewMedicalRecord(models.MedicRecord) (error)
	Report() ([]response.OutPatientReportResponse, error)
	Proceed(models.MedicRecord) (error)
	ListForDoctor(int) ([]response.OutPatientResponse, error)
	ListForNurse(string) ([]response.OutPatientResponse, error)
	FilterByDate(string, string) ([]response.OutPatientResponse, error)
}
