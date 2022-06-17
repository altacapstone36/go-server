package models

import (
	"fmt"
	"go-hospital-server/internal/utils/errors"
	"time"

	"gorm.io/gorm"
)

type MedicRecord struct {
	ID uint `gorm:"primarykey:autoIncrement:false"`
	SerialNumber string
	BloodTension int
	Height int
	Weight int
	BodyTemperature int
	Complaint string
	Diagnose string
	Prescription string
	PatientID uint
	Patient Patient
	MedicalSession MedicalSession `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

func (mr *MedicRecord) BeforeCreate(tx *gorm.DB) (err error) {
	var residence string
	var year int
	var id int64
	ms := mr.MedicalSession
	
	fk := []string{"patients", "medical_staffs", "medical_facilities", "sessions"}
	fk_val := []uint{mr.PatientID, ms.MedicalStaffID, ms.MedicalFacilityID, ms.SessionID}
	fk_msg := []string{"Patient", "Medical Staff", "Medical Facility", "Session"}

	for i, v := range fk {
		var c int64
		tx.Table(v).Where("id = ?", fk_val[i]).Count(&c)
		if c == 0 {
			errMsg := fmt.Sprintf("No %s with ID %d", fk_msg[i], fk_val[i])
			err = errors.New(417, errMsg)
			return
		}
	}



	tx.Model(&mr).Count(&id)
	tx.Model(&mr.Patient).Select("resident_registration").
		Scan(&residence)

	year = time.Now().Year()
	residence = residence[len(residence)-3:]

	mr.SerialNumber = fmt.Sprintf("RM/%s/%d/%05d", residence, year, id+1)

	return
}
