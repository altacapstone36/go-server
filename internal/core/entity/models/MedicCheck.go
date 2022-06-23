package models

import (
	"fmt"
	"go-hospital-server/internal/utils/errors"

	"gorm.io/gorm"
)

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

func (mr *MedicCheck) BeforeCreate(tx *gorm.DB) (err error) {
	var c int64
	
	tx.Debug().Table("users").Where("id = ?", mr.UserID).
		Where("role_id = 3").Count(&c)
	if c == 0 {
		errMsg := fmt.Sprintf("No Nurse with ID #%d available", mr.UserID)
		err = errors.New(417, errMsg)
		return
	}

	tx.Model(&mr).Where("id = ?", mr.ID).
		Where("status = 0").Count(&c)
	if c != 0 {
		errMsg := fmt.Sprintf("Medic Record #%d has unfilled Medical Check", mr.ID)
		err = errors.New(417, errMsg)
		return
	}

	return
}

func (mr *MedicCheck) BeforeUpdate(tx *gorm.DB) (err error) {
	var c int64

	tx.Model(&mr).Where("id = ?", mr.ID).
		Where("status = 1").Count(&c)
	if c != 0 {
		errMsg := fmt.Sprintf("Medical Check #%d already filled", mr.ID)
		err = errors.New(417, errMsg)
		return
	}

	return
}
