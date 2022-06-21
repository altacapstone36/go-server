package repository

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"
)

type OutPatientRepository interface {
	NewMedicalRecord(models.MedicRecord) (error)
	ProceedDoctor(models.MedicRecord) (error)
	AssignNurse(models.MedicCheck) (error)
	ProceedNurse(models.MedicCheck) (error)
	DoctorFindAll(int) ([]response.OutPatientResponse, error)
	NurseFindAll(int) ([]response.OutPatientResponse, error)
	FindByID(int) ([]response.OutPatientResponse, error)
	FindByDate(string, string) ([]response.OutPatientResponse, error)
	Report() ([]response.OutPatientReportResponse, error)
	ReportLog(int, string) ([]response.OutPatientReportLogResponse, error)
}
