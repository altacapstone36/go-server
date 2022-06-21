package database

import (
	"go-hospital-server/internal/core/entity/models"

	"gorm.io/gorm"
)

func migrateDB(db *gorm.DB) (err error) {
	err = db.AutoMigrate(
		models.User{},
		models.MedicalFacility{},
		models.Patient{},
		models.MedicRecord{},
		models.MedicCheck{},
		models.Session{},
		models.Schedule{},
		models.MedicalSession{},
	)
	return
}
