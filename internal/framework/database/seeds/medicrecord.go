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
