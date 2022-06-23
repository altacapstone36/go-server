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
	DoctorFindAll(int) ([]response.OutPatient, error)
	NurseFindAll(int) ([]response.OutPatient, error)
	FindByID(int) ([]response.OutPatientSimple, error)
	FindByDate(string, string) ([]response.OutPatient, error)
	Report() ([]response.OutPatientReport, error)
	ReportLog(int, string) ([]response.OutPatientReportLog, error)
}
