package models

type Schedule struct {
	MedicalStaffID uint `gorm:"primarykey;autoIncrement:false"`
	MedicalStaff MedicalStaff
	MedicalFacilityID uint
	MedicalFacility MedicalFacility
	SessionID uint
	Session Session
	Date string
}
