package models

import (
	"fmt"
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
	
	tx.Model(&mr).Count(&id)
	tx.Model(&mr.Patient).Select("resident_registration").
		Scan(&residence)

	year = time.Now().Year()
	residence = residence[len(residence)-3:]

	mr.SerialNumber = fmt.Sprintf("RM/%s/%d/%05d", residence, year, id+1)

	return
}
