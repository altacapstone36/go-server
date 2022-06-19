package models

import "gorm.io/gorm"

type	MedicalSession struct {
	MedicRecordID uint
	MedicalStaffID uint `json:"medical_staff_id"`
	MedicalFacilityID uint `json:"medical_facility_id"`
	SessionID uint `json:"session_id"`
	DateCheck string `json:"date_check"`
	Queue int
}

func (ms *MedicalSession) BeforeCreate(tx *gorm.DB) (err error) {
	var count int64

	tx.Model(&ms).
		Where("medical_staff_id = ?", ms.MedicalStaffID).
		Where("medical_facility_id = ?", ms.MedicalFacilityID).
		Where("date_check = ?", ms.DateCheck).
		Where("session_id = ?", ms.SessionID).
		Count(&count)

	ms.Queue = int(count)+1

	return
}
