package seeds

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/utils"

	"gorm.io/gorm"
)

func userSeeder() Seed {
	password, _ := utils.HashPassword("password")
	seeds := []models.User{
		{
			Name: "Administrator",
			Email: "admin@web.io",
			Password: password,
			LevelID: 1,
			Status: 1,
		},
		{
			Name: "Ach. Dafid",
			Email: "dafid@web.io",
			Password: password,
			LevelID: 2,
			Status: 1,
		},
		{
			Name: "Syarif Ubaidillah",
			Email: "syarif@web.io",
			Password: password,
			LevelID: 3,
			Status: 1,
		},
	}
	model := &models.User{}

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
