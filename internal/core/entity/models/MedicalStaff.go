package models

type MedicalStaff struct {
	ID int `json:"id" gorm:"primaryKey"`
	FullName string `json:"full_name"`
	Gender string `json:"gender"`
	UserID uint `json:"user_id"`
	User User `json:"user"`
	MedicalFacilityID uint `json:"facility_id"`
	MedicalFacility MedicalFacility `json:"facility"`
}

func (*MedicalStaff) TableName() string {
	return "medical_staff"
}
