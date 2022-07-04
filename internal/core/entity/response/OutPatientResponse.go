package response

type OutPatient struct {
	ID           uint   `json:"id" example:"1"`
	SerialNumber string `json:"serial_number" example:"RM/748/2022/001"`
	FullName     string `json:"full_name" example:"Faizur Ramadhan"`
	Code         string `json:"patient_code" example:"RM0001"`
	Complaint    string `json:"complaint" example:"sakit perut"`
	Doctor       string `json:"doctor" example:"Alsyad Ahmad"`
	Session      string `json:"session" example:"08:00"`
	DateCheck    string `json:"date_check" example:"2022-06-22"`
	Queue        int    `json:"queue" example:"1"`
}

type OutPatientDetails struct {
	ID           int         `json:"id"`
	PatientName  string      `json:"patient_name"`
	NationalID   int         `json:"national_id"`
	SerialNumber string      `json:"serial_number" example:"RM/748/2022/0001"`
	Complaint    string      `json:"complaint" example:"Sakit Perut"`
	Diagnose     string      `json:"diagnose" example:"Maag"`
	Prescription string      `json:"prescription" example:"Entrostop"`
	Doctor       string      `json:"doctor" example:"Alsyad Ahmad"`
	Facility     string      `json:"facility" example:"General"`
	DateCheck    string      `json:"date_check" example:"2022-06-17"`
	MedicCheck   *MedicCheck `json:"medic_check"`
}

type OutPatientReportLog struct {
	SerialNumber string `json:"serial_number" example:"RM/748/2022/001"`
}
