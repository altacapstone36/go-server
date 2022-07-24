package query

import (
	"fmt"
	"strings"
)

func RoleCode() string {
	return "SELECT `code` FROM `roles` WHERE id = ?"
}

func UserRole() string {
	return "SELECT count(*) FROM `users` WHERE role_id = ?"
}

func Duplicate(table, field, value string, id int) string {
		var where []string

		where = append(where, fmt.Sprintf("%s = '%s'", field, value))

		if id != 0 {
			where = append(where, fmt.Sprintf("id != %d", id))
		}

	return "SELECT count(*) FROM `"+ table +"` WHERE " + strings.Join(where, " AND ")
}

func RoleLimit(facility_id int) string {
	var where []string

	where = append(where, "role_id = ?")
	if facility_id != 0 {
		w := fmt.Sprintf("medical_facility_id = %d", facility_id)
		where = append(where, w)
	}
	return "SELECT count(*) FROM `users` WHERE " + strings.Join(where, " AND ")
}

func AvailableFacility() string {
	return "SELECT count(*) FROM `medical_facilities` WHERE id = ?"
}

func AvailablePatient() string {
	return "SELECT count(*) FROM `patients` WHERE code = ?"
}

func PatientMedicRecord() string {
	return "SELECT count(*) FROM `medic_records` WHERE patient_code = ?"
}

func StaffCheck(id int) string {
	return fmt.Sprintf("SELECT count(*) FROM `users` WHERE code = ? AND medical_facility_id = %d", id)
}
