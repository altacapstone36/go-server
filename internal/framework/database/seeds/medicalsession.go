package seeds

import (
	"go-hospital-server/internal/core/entity/models"

	"gorm.io/gorm"
)

func medicalSessionSeeder() Seed {
	seeds := []models.MedicalSession{
		{
			MedicRecordID:     1,
			MedicalFacilityID: 1,
			SessionID:         1,
			DateCheck:         "2022-07-17",
			Queue:             1,
		},
	}
	model := &models.MedicalSession{}

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
