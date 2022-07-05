package check

import (
	"go-hospital-server/internal/utils/errors"

	"gorm.io/gorm"
)

type DBAction string
const DELETE DBAction = "delete"
const FIND DBAction = "find"
const UPDATE DBAction = "update"
const CREATE DBAction = "create"

// check database error
// with parameter *gorm.DB and DBAction
// example:
// DB := *gorm.DB
// db := DB.Find(models.User{})
// err := check.DBRecord(db, check.FIND)
// DBAction available in DELETE, UPDATE, FIND.
func DBRecord(db *gorm.DB, action DBAction) (err error) {
	if db.Error != nil {
		err = errors.ErrInternalServer
		return
	}

	if db.RowsAffected != 0 {
		return
	}

	if action == DELETE || action == FIND {
		err = errors.ErrNoRecord
	} else if action == UPDATE {
		err = errors.ErrNoChange
	}

	return
}
