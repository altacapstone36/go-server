package response

type OutPatient struct {
	ID uint `json:"id" example:"1"`
	SerialNumber string `json:"serial_number" example:"RM/748/2022/001"`
	FullName string `json:"full_name" example:"Faizur Ramadhan"`
	Code string `json:"code" example:"RM0001"`
	Complaint string `json:"complaint" example:"sakit perut"`
	Doctor string `json:"medical_staff" example:"Alsyad Ahmad"`
	TimeStart string `json:"session" example:"08:00"`
	DateCheck string `json:"date_check" example:"2022-06-22"`
	Queue int `json:"queue" example:"1"`
}

type OutPatientSimple struct {
	ID uint `json:"id" example:"1"`
	SerialNumber string `json:"serial_number" example:"RM/748/2022/001"`
	FullName string `json:"full_name" example:"Faizur Ramadhan"`
	Code string `json:"code" example:"RM0001"`
}

type OutPatientReport struct {
	ID uint `json:"id" example:"1"`
	SerialNumber string `json:"serial_number" example:"RM/748/2022/001"`
	Code string `json:"code" example:"RM0001"`
	ResidentRegistration string `json:"resident_registration" example:"8729301745162748"`
	FullName string `json:"full_name" example:"Faizur Ramadhan"`
	Complaint string `json:"complaint" example:"sakit perut"`
	Diagnose string `json:"diagnose" example:"maag"`
	Prescription string `json:"prescription" example:"entrostop"`
	Doctor string `json:"doctor" example:"Alsyad Ahmad"`
	Name string `json:"name" example:"General"`
	DateCheck string `json:"date_check" example:"2022-06-22"`
}

type OutPatientReportLog struct {
	SerialNumber string `json:"serial_number" example:"RM/748/2022/001"`
}
