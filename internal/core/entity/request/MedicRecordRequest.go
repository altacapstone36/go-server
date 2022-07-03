package request

import (
	vs "go-hospital-server/internal/utils/validators"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type AdminMedicRecord struct {
	PatientCode       string `json:"patient_code" example:"1"`
	UserCode          string `json:"doctor_code" example:"DR00002"`
	MedicalFacilityID uint   `json:"facility_id" example:"1"`
	Complaint         string `json:"complaint" example:"sakit perut"`
	SessionID         uint   `json:"session_id" example:"1"`
	DateCheck         string `json:"date_check" example:"2022-06-24"`
}

type DoctorMedicRecord struct {
	ID           int    `json:"-"`
	Diagnose     string `json:"diagnose" example:"maag"`
	Prescription string `json:"prescription" example:"entrostop"`
}

type NurseMedicRecord struct {
	ID              int `json:"-"`
	BloodTension    int `json:"blood_tension" example:"122"`
	Height          int `json:"height" example:"55"`
	Weight          int `json:"weight" example:"165"`
	BodyTemperature int `json:"body_temp" example:"31"`
}

type AssignNurse struct {
	ID   int    `json:"-"`
	Code string `json:"nurse_code" example:"NR00003"`
}

func (s AdminMedicRecord) Validate() (err error) {
	err = v.ValidateStruct(&s,
		v.Field(&s.PatientCode, v.Required, v.By(vs.AvailablePatient), v.By(vs.PatientMedicRecord)),
		v.Field(&s.MedicalFacilityID, v.NilOrNotEmpty, v.By(vs.AvailableFacility)),
		v.Field(&s.UserCode, v.Required, v.By(vs.StaffCheck(int(s.MedicalFacilityID))), v.By(vs.CodeCheck("DCR"))),
		v.Field(&s.SessionID, v.NilOrNotEmpty),
		v.Field(&s.DateCheck, v.Required, v.Date("2006-01-02").Error("date value must use format yyyy-mm-dd")),
	)
	return
}

func (s DoctorMedicRecord) Validate() (err error) {
	err = v.ValidateStruct(&s,
		v.Field(&s.ID, v.By(vs.ProcessDoctor), v.By(vs.MedicRecordCheck)),
		v.Field(&s.Diagnose, v.Required, is.Alpha),
		v.Field(&s.Prescription, v.Required),
	)
	return
}

func (s NurseMedicRecord) Validate() (err error) {
	err = v.ValidateStruct(&s,
		v.Field(&s.ID, v.By(vs.ProcessNurse)),
		v.Field(&s.BodyTemperature, v.Required),
		v.Field(&s.Height, v.Required),
		v.Field(&s.Weight, v.Required),
		v.Field(&s.BloodTension, v.Required),
	)
	return
}

func (s AssignNurse) Validate() (err error) {
	err = v.ValidateStruct(&s,
		v.Field(&s.ID, v.NilOrNotEmpty, v.By(vs.NurseAssign)),
		v.Field(&s.Code, v.Required, v.By(vs.CodeCheck("NRS")), v.By(vs.NurseFacility(s.ID))),
	)
	return
}
