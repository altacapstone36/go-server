package seeds

import (
	"go-hospital-server/internal/core/entity/models"
	"time"

	"gorm.io/gorm"
)

func patientSeeder() Seed {
	seeds := []models.Patient{
		{
			ResidentRegistration: 8729301745162748,
			FullName: "Faizur Ramadhan",
			Address: "Sumenep",
			Gender: "Male",
			BirthDate: time.Date(2000, time.April, 20, 0, 0, 0, 0, &time.Location{}),
			BloodType: "A",
		},
		{
			ResidentRegistration: 8761435451636421,
			FullName: "Ach. Novil",
			Address: "Sumenep",
			Gender: "Male",
			BirthDate: time.Date(2000, time.March, 10, 0, 0, 0, 0, &time.Location{}),
			BloodType: "A",
		},
	}
	model := &models.Patient{}

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
