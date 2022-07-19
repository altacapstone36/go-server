package models

type Schedule struct {
	UserCode          string `gorm:"index"`
	MedicalFacilityID uint
	SessionID         uint
	Date              string

	// User User `gorm:"foreignkey:UserCode;references:code"`
}
