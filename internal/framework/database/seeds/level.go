package seeds

import (
	"go-hospital-server/internal/core/entity/models"

	"gorm.io/gorm"
)

func levelSeeder() Seed {
	seeds := []models.Role{
		{Name: "Admin", Code: "ADM"},
		{Name: "Doctor", Code: "DCR"},
		{Name: "Nurse", Code: "NRS"},
	}
	model := &models.Role{}

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
