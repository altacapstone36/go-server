package validators

import (
	"errors"
	"fmt"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

var _DB *gorm.DB

func NewValidator(db *gorm.DB) {
	_DB = db
}

// Check data duplication
// Param table string
// Param field string
// Param id int
func Duplicate(table, field string, id int) validation.RuleFunc {
	return func(value interface{}) (err error) {
		var c int64
		var where []string
		field_name := strings.Split(field, "_")

		where = append(where, fmt.Sprintf("%s = '%s'", field, value))

		if id != 0 {
			where = append(where, fmt.Sprintf("id != %d", id))
		}

		_DB.Table(table).
			Where(strings.Join(where, " AND ")).Count(&c)

		if c != 0 {
			msg := fmt.Sprintf("duplicate %s %s.", strings.Join(field_name, " "), value.(string))
			err = errors.New(msg)
		}

		return
	}
}

func Email(id int) validation.RuleFunc {
	return func(value interface{}) (err error) {
		var c int64
		var where string

		if id != 0 {
			where = fmt.Sprintf("id != %d", id)
		}

		_DB.Table("users").
			Where("email = ?", value).
			Where(where).Count(&c)

		if c != 0 {
			msg := fmt.Sprintf("duplicate email %s.", value.(string))
			err = errors.New(msg)
		}

		return
	}
}

func AvailableFacility(value interface{}) (err error) {
	var c int64

	_DB.Table("medical_facilities").
		Where("id = ?", value).
		Count(&c)

	if c == 0 {
		msg := fmt.Sprintf("no facility with id #%d.", value)
		err = errors.New(msg)
	}

	return
}

func AvailablePatient(value interface{}) (err error) {
	var c int64

	_DB.Table("patients").
		Where("code = ?", value).
		Count(&c)

	if c == 0 {
		msg := fmt.Sprintf("no patient with code %s.", value.(string))
		err = errors.New(msg)
	}

	return
}

func PatientMedicRecord(value interface{}) (err error) {
	var c int64

	_DB.Table("medic_records").
		Where("patient_code = ?", value).
		Count(&c)

	if c != 0 {
		msg := fmt.Sprintf("patient %s have %d unfilled medic record.", value.(string), c)
		err = errors.New(msg)
	}

	return
}

func MedicRecordCheck(value interface{}) (err error) {
	var c int64

	_DB.Table("medic_checks").
		Where("medic_record_id = ?", value).
		Where("status = 0").
		Count(&c)

	if c != 0 {
		msg := fmt.Sprintf("medic record %d have %d unfilled medic check.", value.(int), c)
		err = errors.New(msg)
	}

	return
}

func RoleLimit(facility_id int) validation.RuleFunc {
	return func(value interface{}) (err error) {
		var c int64
		var where string

		if facility_id != 0 {
			where = fmt.Sprintf("medical_facility_id = %d", facility_id)
		}

		_DB.Table("users").
			Where("role_id = ?", value).
			Where(where).Count(&c)

		if c >= 2 && value == 3 {
			msg := fmt.Sprintf("nurse in facility #%d exceed", facility_id)
			err = errors.New(msg)
		}

		return
	}
}

func StaffCheck(facility_id int) validation.RuleFunc {
	return func(value interface{}) (err error) {
		var c int64
		var where string

		if facility_id != 0 {
			where = fmt.Sprintf("medical_facility_id = %d", facility_id)
		}

		_DB.Table("users").
			Where("code = ?", value).
			Where(where).Count(&c)

		if c == 0 {
			msg := fmt.Sprintf("no staff with code %s in facility #%d", value, facility_id)
			err = errors.New(msg)
		}

		return
	}
}

func ProcessDoctor(value interface{}) (err error) {
	var c int64

	_DB.Table("medic_records").
		Where("id = ?", value).
		Where("status = 1").
		Count(&c)

	if c != 0 {
		msg := fmt.Sprintf("medic record #%d filled", value)
		err = errors.New(msg)
	}

	return
}

func ProcessNurse(value interface{}) (err error) {
	var c int64

	_DB.Table("medic_checks").
		Where("medic_record_id = ?", value).
		Where("status = 1").
		Count(&c)

	if c != 0 {
		msg := fmt.Sprintf("medic check from record #%d filled", value)
		err = errors.New(msg)
	}

	return
}

func NurseAssign(value interface{}) (err error) {
	var c int64

	_DB.Table("medic_checks").
		Where("medic_record_id = ?", value).
		Count(&c)

	if c != 0 {
		msg := fmt.Sprintf("medic check from record #%d assigned", value)
		err = errors.New(msg)
	}

	return
}

func NurseFacility(id int) validation.RuleFunc {
	return func(value interface{}) (err error) {
		var c, facility_id int64

		_DB.Table("medical_sessions").Select("medical_facility_id").
			Where("medic_record_id = ?", id).
			Scan(&facility_id)

		_DB.Table("users").
			Where("code = ?", value).
			Where("medical_facility_id = ?", facility_id).
			Count(&c)

		if c == 0 {
			msg := fmt.Sprintf("no nurse with code %s in facility #%d", value, facility_id)
			err = errors.New(msg)
		}

		return
	}
}

func CodeCheck(code string) validation.RuleFunc {
	return func(value interface{}) (err error) {
		var c int64

		_DB.Table("users").
			Where("code LIKE ?", code+"%").
			Where("code = ?", value).
			Count(&c)

		if c == 0 {
			msg := fmt.Sprintf("only code %s to assigned", code)
			err = errors.New(msg)
		}

		return
	}
}
