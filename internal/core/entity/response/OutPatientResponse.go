package response

type OutPatientResponse struct {
	ID uint `json:"id"`
	SerialNumber string `json:"serial_number"`
	FullName string `json:"full_name"`
	Code string `json:"code"`
	Complaint string `json:"complaint"`
	Doctor string `json:"doctor"`
	TimeStart string `json:"session"`
	DateCheck string `json:"date_check"`
	Queue int `json:"queue"`
}

type OutPatientReportResponse struct {
	ID uint `json:"id"`
	SerialNumber string `json:"serial_number"`
	Code string `json:"code"`
	ResidentRegistration string `json:"resident_registration"`
	FullName string `json:"full_name" sql:"unique"`
	BloodType string `json:"blood_type"`
	BloodTension int `json:"blood_tension"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	BodyTemperature int `json:"body_temp"`
	Complaint string `json:"complaint"`
	Diagnose string `json:"diagnose"`
	Prescription string `json:"prescription"`
	Doctor string `json:"doctor"`
	Facility string `json:"facility"`
	DateCheck string `json:"date_check"`
}

type OutPatientReportLogResponse struct {
	SerialNumber string `json:"serial_number"`
}
