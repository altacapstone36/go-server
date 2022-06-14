package models

import (
	"fmt"
	"go-hospital-server/internal/utils/errors"

	"gorm.io/gorm"
)

type MedicalStaff struct {
	ID uint `json:"id" gorm:"primaryKey"`
	FullName string `json:"full_name"`
	Gender string `json:"gender"`
	UserID uint `json:"user_id"`
	User User `json:"user"`
	MedicalFacilityID uint `json:"facility_id"`
	MedicalFacility MedicalFacility `json:"facility"`
}

func (ms *MedicalStaff) BeforeCreate(tx *gorm.DB) (err error) {

	if err = tx.Model(&ms).Where("full_name = ?", ms.FullName).Error;
	err == nil {
		msg := fmt.Sprintf("duplicate name for %s found, please use another name", ms.FullName)
		err = errors.New(203, msg)
		return
	}

	return
}

