package models

type MedicalFacility struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (*MedicalFacility) TableName() string {
	return "medical_facility"
}
