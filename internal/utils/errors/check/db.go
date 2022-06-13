package check

import (
	e "errors"
	"go-hospital-server/internal/utils"
	"go-hospital-server/internal/utils/errors"
	"go-hospital-server/internal/utils/logger"

	"gorm.io/gorm"
)


func Record(res interface{}, err error) error {
	if utils.IsEmpty(res) && res != nil || e.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(404, "record not found")
	}

	if err != nil {
		logger.WriteLog(err)
		return errors.New(500, "internal server error")
	}

	return nil
}
