package models

type Schedule struct {
	UserCode          string `json:"doctor_code" gorm:"index"`
	SessionID         uint	`json:"session_id"`
	Date              string `json:"date_check"`
}
