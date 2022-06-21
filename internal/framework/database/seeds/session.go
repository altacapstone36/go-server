package seeds

import (
	"go-hospital-server/internal/core/entity/models"

	"gorm.io/gorm"
)

func sessionSeeder() Seed {
	seeds := []models.Session{
		{
			Name:      "Pagi",
			TimeStart: "08:00",
			TimeEnd:   "11:00",
		},
		{
			Name:      "Siang",
			TimeStart: "13:00",
			TimeEnd:   "15:00",
		},
		{
			Name:      "Sore",
			TimeStart: "16:00",
			TimeEnd:   "18:00",
		},
		{
			Name:      "Malam",
			TimeStart: "19:00",
			TimeEnd:   "22:00",
		},
	}
	model := &models.Session{}

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
