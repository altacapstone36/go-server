package request

import (
	vs "go-hospital-server/internal/utils/validators"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

type DoctorSchedule struct {
	MedicalFacilityID uint   `json:"facility_id"`
	UserCode          string `json:"doctor_code"`
	DateCheck         string `json:"date_check"`
	SessionID         uint   `json:"session_id"`
}

func (ds DoctorSchedule) Validate() (err error) {
	err = v.ValidateStruct(&ds,
		v.Field(&ds.MedicalFacilityID, v.Required, v.NilOrNotEmpty, v.By(vs.AvailableFacility)),
		v.Field(&ds.UserCode, v.Required, v.By(vs.StaffCheck(int(ds.MedicalFacilityID))), v.By(vs.CodeCheck("DCR"))),
		v.Field(&ds.DateCheck, v.Required, v.Date("2006-01-02").Error("date value must use format yyyy-mm-dd")),
		v.Field(&ds.SessionID, v.NilOrNotEmpty),
	)
	return
}
