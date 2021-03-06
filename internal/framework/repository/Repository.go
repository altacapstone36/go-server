package repository

import (
	"go-hospital-server/internal/core/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB, mongo *mongo.Database) *repository.Repository {
	return &repository.Repository{
		Auth:       NewAuthRepository(db, mongo),
		Patient:    NewPatientRepository(db),
		OutPatient: NewOutPatientRepository(db),
		User:       NewUserRepository(db),
		Facility:   NewFacilityRepository(db),
		Schedule:   NewScheduleRepository(db),
		Session:    NewSessionRepository(db),
	}
}
