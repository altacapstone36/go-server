package models

import "gorm.io/gorm"

type MedicalSession struct {
	MedicRecordID     uint   `gorm:"primarykey;autoincrement:false"`
	UserCode          string `json:"doctor_code" gorm:"index"`
	MedicalFacilityID uint   `json:"facility_id"`
	SessionID         uint   `json:"session_id"`
	DateCheck         string `json:"date_check"`
	Queue             int

	MedicalFacility MedicalFacility
	MedicRecord     *MedicRecord
	Session         Session
	User            User `gorm:"foreignkey:UserCode;references:code"`
}

func (ms *MedicalSession) BeforeCreate(tx *gorm.DB) (err error) {
	var count int64

	tx.Model(&ms).
		Where("user_code = ?", ms.UserCode).
		Where("medical_facility_id = ?", ms.MedicalFacilityID).
		Where("date_check = ?", ms.DateCheck).
		Where("session_id = ?", ms.SessionID).
		Count(&count)

	ms.Queue = int(count) + 1

	return
}
