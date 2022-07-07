package models

type DoctorSchedule struct {
	MedicalFacilityID uint   `json:"facility_id"`
	UserCode          string `json:"doctor_code" gorm:"index"`
	DateCheck         string `json:"date_check"`
	SessionID         uint   `json:"session_id"`

	MedicalFacility MedicalFacility
	Session         Session
	User            User `gorm:"foreignkey:UserCode;references:code"`
}
