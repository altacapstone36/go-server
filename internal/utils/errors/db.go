package errors

import (
	"errors"
	"go-hospital-server/internal/utils"
	"go-hospital-server/internal/utils/logger"

	"gorm.io/gorm"
)


func CheckError(res interface{}, err error) error {
	if utils.IsEmpty(res) && res != nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return New(404, "record not found")
	}

	if err != nil {
		logger.WriteLog(err)
		return New(500, "internal server error")
	}

	return nil
}
