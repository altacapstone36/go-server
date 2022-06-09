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
			Email: "alsyadahmad@holyhos.co.id",
			Password: password,
			RoleID: 2,
			Status: 1,
		},
		{
			Email: "priscillahalim@holyhos.co.id",
			Password: password,
			RoleID: 3,
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
