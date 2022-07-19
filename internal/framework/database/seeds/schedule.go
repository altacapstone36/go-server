package seeds

import (
	"go-hospital-server/internal/core/entity/models"

	"gorm.io/gorm"
)

func scheduleSeeder() Seed {
	seeds := []models.Schedule{
		{
			UserCode:          "DCR00001",
			SessionID:         1,
			Date:              "2022-06-17",
		},
		{
			UserCode:          "DCR00001",
			SessionID:         1,
			Date:              "2022-06-18",
		},
	}
	model := &models.Schedule{}

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
