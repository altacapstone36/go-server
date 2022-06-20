package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type AdminMedicRecord  struct {
	PatientID uint `json:"patient_id" example:"1"`
	MedicalStaffID uint `json:"medical_staff_id" example:"1"`
	MedicalFacilityID uint `json:"medical_facility_id" example:"1"`
	Complaint string `json:"complaint" example:"sakit perut"`
	SessionID uint `json:"session_id" example:"1"`
	DateCheck string `json:"date_check" example:"2022-06-24"`
}

type DoctorMedicRequest struct {
	ID uint `json:"id" example:"1"`
	Diagnose string `json:"diagnose" example:"maag"`
	Prescription string `json:"prescription" example:"entrostop"`
}

type NurseMedicRequest struct {
	ID uint `json:"id"`
	BloodTension int `json:"blood_tension" example:"122"`
	Height int `json:"height" example:"55"`
	Weight int `json:"weight" example:"165"`
	BodyTemperature int `json:"body_temp" example:"31"`
}

func (s AdminMedicRecord) Validate() (err error) {
	err = validation.ValidateStruct(&s,
		validation.Field(&s.MedicalFacilityID, validation.NilOrNotEmpty),
		validation.Field(&s.MedicalStaffID, validation.NilOrNotEmpty),
		validation.Field(&s.SessionID, validation.NilOrNotEmpty),
		validation.Field(&s.DateCheck, validation.Required, validation.Date("2006-01-02")),
	)
	return
}

func (s DoctorMedicRequest) Validate() (err error) {
	err = validation.ValidateStruct(&s,
		validation.Field(&s.Diagnose, validation.Required, is.Alpha),
		validation.Field(&s.Prescription, validation.Required),
		validation.Field(&s.ID, validation.Required),
	)
	return
}

func (s NurseMedicRequest) Validate() (err error) {
	err = validation.ValidateStruct(&s,
		validation.Field(&s.ID, validation.Required),
		validation.Field(&s.BodyTemperature, validation.Required),
		validation.Field(&s.Height, validation.Required),
		validation.Field(&s.Weight, validation.Required),
		validation.Field(&s.BloodTension, validation.Required),
	)
	return
}
