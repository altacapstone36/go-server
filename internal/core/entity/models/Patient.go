package models

import "time"

type Patient struct {
	ID uint `json:"id" gorm:"primaryKey"`
	ResidentRegistration int64 `json:"resident_registration"`
	FullName string `json:"full_name"`
	Address string `json:"address"`
	Gender string `json:"gender"`
	BirthDate time.Time`json:"birthdate"`
	BloodType string `json:"blood_type"`
}
