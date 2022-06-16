package models

import "gorm.io/gorm"

type	MedicalSession struct {
	MedicRecordID uint
	MedicalFacilityID uint
	MedicalStaffID uint
	SessionID uint
	DateCheck string
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
