package repository

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"
)

type FacilityRepository interface {
	FindAll() ([]response.Facility, error)
	FindByID(int) (response.FacilityDetails, error)
	Create(models.MedicalFacility) error
	Update(models.MedicalFacility) error
	Delete(int) error
}
