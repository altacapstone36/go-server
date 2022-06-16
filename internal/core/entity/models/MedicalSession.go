package models

type	MedicalSession struct {
	MedicRecordID uint `gorm:"primaryKey;autoIncrement:false"`
	MedicRecord MedicRecord
	MedicalFacilityID uint
	MedicalFacility MedicalFacility
	SessionID uint
	Session Session
	DateCheck string
	Queue int
}
