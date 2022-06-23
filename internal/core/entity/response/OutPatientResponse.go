package response

type OutPatient struct {
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

type OutPatientSimple struct {
	ID uint `json:"id"`
	SerialNumber string `json:"serial_number"`
	FullName string `json:"full_name"`
	Code string `json:"code"`
}

type OutPatientReport struct {
	ID uint `json:"id"`
	SerialNumber string `json:"serial_number"`
	Code string `json:"code"`
	ResidentRegistration string `json:"resident_registration" example:"8729301745162748"`
	FullName string `json:"full_name"`
	Complaint string `json:"complaint"`
	Diagnose string `json:"diagnose"`
	Prescription string `json:"prescription"`
	Doctor string `json:"doctor"`
	Facility string `json:"facility"`
	DateCheck string `json:"date_check"`
}

type OutPatientReportLog struct {
	SerialNumber string `json:"serial_number"`
}
