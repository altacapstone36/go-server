package models

import (
	"fmt"
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

	tx.Model(&mr).Count(&id)
	tx.Model(&mr.Patient).Select("resident_registration").
		Scan(&residence)

	year = time.Now().Year()
	residence = residence[len(residence)-3:]

	mr.SerialNumber = fmt.Sprintf("RM/%s/%d/%05d", residence, year, id+1)

	return
}
