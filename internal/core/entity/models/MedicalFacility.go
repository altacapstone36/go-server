package models

type MedicalFacility struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
