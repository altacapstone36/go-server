package models

type MedicRecord struct {
	ID uint `gorm:"primaryKey"`
	// SerialNumber uint
	BloodTension int
	Height int
	Weight int
	BodyTemperature int
	Complaint string
	Diagnose string
	Prescription string
	PatientID uint
	Patient Patient
	MedicalStaffID uint
	MedicalStaff MedicalStaff
}
