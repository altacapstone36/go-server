package repository

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"
	"time"
)

type OutPatientRepository interface {
	NewMedicalRecord(models.MedicRecord) (error)
	Proceed(models.MedicRecord) (error)
	ListAvailable(int, string) ([]response.OutPatientResponse, error)
	FilterByDate(time.Time, time.Time) ([]response.OutPatientResponse, error)
}
