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
	DoctorFindAll(string) ([]response.OutPatient, error)
	NurseFindAll(string) ([]response.OutPatient, error)
	FindByID(int) (response.OutPatientDetails, error)
	FindByDate(string, string) ([]response.OutPatient, error)
	Report() ([]response.OutPatientDetails, error)
	ReportLog(string, string) ([]response.OutPatientReportLog, error)
}
