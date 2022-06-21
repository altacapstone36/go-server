package seeds

import (
	"go-hospital-server/internal/core/entity/models"

	"gorm.io/gorm"
)

func medicRecordSeeder() Seed {
	seeds := []models.MedicRecord{
		{
			Complaint:       "Sakit Telinga",
			PatientID:       1,
			MedicCheck: models.MedicCheck{
				UserID: 3,
			},
			MedicalSession: models.MedicalSession{
				MedicalFacilityID: 1,
				UserID: 2,
				SessionID:         1,
				DateCheck:         "2022-07-17",
			},
		},
	}
	model := &models.MedicRecord{}

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
