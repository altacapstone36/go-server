package models

import (
	"fmt"
	"go-hospital-server/internal/utils/errors"

	"gorm.io/gorm"
)

type MedicalFacility struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (mf *MedicalFacility) BeforeCreate(tx *gorm.DB) (err error) {
	var c int64

	tx.Model(&mf).Select("name").
		Where("name = ?", mf.Name).
		Count(&c);
	if c == 0 {
		errMsg := fmt.Sprintf("Medical Facility with name %s", *&mf.Name)
		err = errors.New(417, errMsg)
		return
	}

	return
}
