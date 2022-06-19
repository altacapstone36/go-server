package seeds

import (
	"go-hospital-server/internal/core/entity/models"

	"gorm.io/gorm"
)

func medicalStaffSeeder() Seed {
	seeds := []models.MedicalStaff{
		{
			FullName: "Rimuru Tempest",
			Gender: "Male",
			UserID: 1,
			MedicalFacilityID: 1,
		},
		{
			FullName: "Alsyad Ahmad",
			Gender: "Male",
			UserID: 2,
			MedicalFacilityID: 1,
		},
		{
			FullName: "Priscilla Halim",
			Gender: "Female",
			UserID: 3,
			MedicalFacilityID: 1,
		},
	}
	model := &models.MedicalStaff{}

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
