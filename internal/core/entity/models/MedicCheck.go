package models

type MedicCheck struct {
	ID uint ` json:"id" gorm:"primarykey:autoIncrement:false"`
	BloodTension int `json:"blood_tension"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	BodyTemperature int `json:"body_temp"`
	MedicRecordID uint `json:"medic_record_id"`
	UserID uint `json:"user_id"`
	Status int
}
