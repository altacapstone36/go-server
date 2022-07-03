package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type MedicRecord struct {
	ID             uint ` json:"id" gorm:"primary_key"`
	SerialNumber   string
	Complaint      string `json:"complaint"`
	Diagnose       string `json:"diagnose"`
	Prescription   string `json:"prescription"`
	PatientCode    string `json:"patient_code" gorm:"index"`
	Status         int
	Patient        Patient        `gorm:"foreignkey:PatientCode;references:code"`
	MedicCheck     MedicCheck     `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	MedicalSession MedicalSession `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

func (mr *MedicRecord) BeforeCreate(tx *gorm.DB) (err error) {
	var residence string
	var year int
	var id int64

	tx.Model(&mr).Count(&id)
	tx.Model(&mr.Patient).Select("national_id").
		Scan(&residence)

	year = time.Now().Year()
	residence = residence[len(residence)-3:]

	mr.SerialNumber = fmt.Sprintf("RM/%s/%d/%05d", residence, year, id+1)

	return
}
