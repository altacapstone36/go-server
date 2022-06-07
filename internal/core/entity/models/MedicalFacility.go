package models

type MedicalFacility struct {
	ID int `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (*MedicalFacility) TableName() string {
	return "medical_facility"
}
