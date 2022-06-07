package seeds

import (
	"go-hospital-server/internal/core/entity/models"

	"gorm.io/gorm"
)

func medicalFacilitySeeder() Seed {
	seeds := []models.MedicalFacility{
		{Name: "General"},
		{Name: "Pediatrician"},
		{Name: "Dentist"},
	}
	model := &models.MedicalFacility{}

	return Seed{
		models: model,
		run: func(db *gorm.DB) (err error) {
			for _, v := range seeds {
				err = db.Create(&v).Error
			}
			return
		},
	}
}
