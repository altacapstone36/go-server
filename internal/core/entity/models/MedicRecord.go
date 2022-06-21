package models

import (
	"fmt"
	"go-hospital-server/internal/utils/errors"
	"time"

	"gorm.io/gorm"
)

type MedicRecord struct {
	ID uint ` json:"id" gorm:"primarykey:autoIncrement:false"`
	SerialNumber string
	Complaint string `json:"complaint"`
	Diagnose string `json:"diagnose"`
	Prescription string `json:"prescription"`
	PatientID uint `json:"patient_id"`
	Patient Patient
	Status int
	MedicCheck MedicCheck `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	MedicalSession MedicalSession `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

func (mr *MedicRecord) BeforeCreate(tx *gorm.DB) (err error) {
	var residence string
	var year int
	var id int64
	var c int64
	ms := mr.MedicalSession
	
	fk := []string{"patients", "users", "medical_facilities", "sessions"}
	fk_val := []uint{mr.PatientID, ms.UserID, ms.MedicalFacilityID, ms.SessionID}
	fk_msg := []string{"Patient", "Medical Staff", "Medical Facility", "Session"}

	for i, v := range fk {
		tx.Table(v).Where("id = ?", fk_val[i]).Count(&c)
		if c == 0 {
			errMsg := fmt.Sprintf("No %s with ID %d", fk_msg[i], fk_val[i])
			err = errors.New(417, errMsg)
			return
		}
	}

	tx.Table(fk[1]).Where("id = ?", fk_val[1])
	if c == 0 {
		errMsg := fmt.Sprintf("Medical Staff #%d not in facility #%d",
			mr.MedicalSession.UserID, mr.MedicalSession.MedicalFacilityID)
		err = errors.New(417, errMsg)
		return
	}
	
	tx.Model(&mr).Count(&id)
	tx.Model(&mr.Patient).Select("resident_registration").
		Scan(&residence)

	year = time.Now().Year()
	residence = residence[len(residence)-3:]

	mr.SerialNumber = fmt.Sprintf("RM/%s/%d/%05d", residence, year, id+1)

	return
}

func (mr *MedicRecord) BeforeUpdate(tx *gorm.DB) (err error) {
	var c int64

	tx.Model(&mr).Select("patient_id").Where("id = ?", mr.ID).Scan(&mr.PatientID)

	tx.Model(&mr).Where("id = ?", mr.ID).
		Where("status = 1").Count(&c)
	if c != 0 {
		errMsg := fmt.Sprintf("Medic Record #%d already filled", mr.ID)
		err = errors.New(417, errMsg)
		return
	}

	tx.Model(&mr.MedicCheck).Where("medic_record_id = ?", mr.ID).
		Where("status = 0").Count(&c)
	if c != 0 {
		errMsg := fmt.Sprintf("Patient #%d have %d unfilled Medical Check", mr.PatientID, c)
		err = errors.New(417, errMsg)
		return
	}

	return
}
