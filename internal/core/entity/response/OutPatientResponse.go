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
