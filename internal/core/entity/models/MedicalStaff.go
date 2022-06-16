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
	Schedule []Schedule `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

func (ms *MedicalStaff) BeforeCreate(tx *gorm.DB) (err error) {
	var name string

	tx.Model(&ms).Select("full_name").
		Where("full_name = ?", ms.FullName).
		Scan(&name);

	if name != "" {
		msg := fmt.Sprintf("duplicate name for %s found, please use another name", ms.FullName)
		err = errors.New(203, msg)
		return
	}

	return
}

