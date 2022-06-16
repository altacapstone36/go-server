package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type AdminMedicRecord  struct {
	PatientID uint `json:"patient_id"`
	MedicalStaffID uint `json:"medical_staff_id"`
	MedicalFacilityID uint `json:"medical_facility_id"`
	Complaint string `json:"complaint"`
	SessionID uint `json:"session_id"`
	DateCheck string `json:"date_check"`
}

type DoctorMedicRequest struct {
	ID uint `json:"id"`
	Diagnose string `json:"diagnose"`
	Prescription string `json:"prescription"`
}

type NurseMedicRequest struct {
	ID uint `json:"id"`
	BloodTension int `json:"blood_tesion"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	BodyTemperature int `json:"body_temp"`
}

func (s AdminMedicRecord) Validate() (err error) {
	err = validation.ValidateStruct(&s,
		validation.Field(&s.MedicalFacilityID, validation.Required, is.Digit),
		validation.Field(&s.MedicalStaffID, validation.Required, is.Digit),
		validation.Field(&s.SessionID, validation.Required, is.Digit),
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
		validation.Field(&s.BloodTension, validation.Required, is.Digit),
		validation.Field(&s.Height, validation.Required, is.Digit),
		validation.Field(&s.Weight, validation.Required, is.Digit),
		validation.Field(&s.BloodTension, validation.Required, is.Digit),
	)
	return
}