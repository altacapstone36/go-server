package models

type MedicCheck struct {
	BloodTension    int    `json:"blood_tension"`
	Height          int    `json:"height"`
	Weight          int    `json:"weight"`
	BodyTemperature int    `json:"body_temp"`
	MedicRecordID   uint   `gorm:"primarykey;autoincrement:false"`
	UserCode        string `json:"nurse_code" gorm:"index"`
	Status          int

	User        User `gorm:"foreignkey:UserCode;references:code"`
	MedicRecord *MedicRecord
}
