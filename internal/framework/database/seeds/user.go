package seeds

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/utils"

	"gorm.io/gorm"
)

func userSeeder() Seed {
	var id uint = 1
	password, _ := utils.HashPassword("password")
	seeds := []models.User{
		{
			FullName: "Rimuru Tempest",
			Email: "rimurutempest@holyhos.co.id",
			Gender: "Male",
			Password: password,
			RoleID: 1,
			Status: 1,
		},
		{
			FullName: "Alsyad Ahmad",
			Email: "alsyadahmad@holyhos.co.id",
			Password: password,
			Gender: "Male",
			MedicalFacilityID: &id,
			RoleID: 2,
			Status: 1,
		},
		{
			FullName: "Priscilla Halim",
			Email: "priscillahalim@holyhos.co.id",
			Password: password,
			Gender: "Female",
			MedicalFacilityID: &id,
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
