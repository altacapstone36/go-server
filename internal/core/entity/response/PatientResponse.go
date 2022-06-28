package response


type Patient struct {
	ID uint `json:"id" example:"1"`
	Code string `json:"code" example:"RM0001"`
	ResidentRegistration string `json:"resident_registration" example:"8729301745162748"`
	FullName string `json:"full_name" sql:"unique" example:"Faizur Ramadhan"`
	Gender string `json:"gender" example:"Male"`
	BirthDate string `json:"birthdate" example:"2001-04-14"`
}

type PatientDetails struct {
	ID uint `json:"id" example:"1"`
	Code string `json:"code" example:"RM0001"`
	ResidentRegistration string `json:"resident_registration" example:"8729301745162748"`
	FullName string `json:"full_name" sql:"unique" example:"Faizur Ramadhan"`
	Address string `json:"address" example:"Sumenep"`
	Gender string `json:"gender" example:"Male"`
	BirthDate string `json:"birthdate" example:"2001-04-14"`
	BloodType string `json:"blood_type" example:"A"`
	MedicRecord []MedicRecord `json:"medic_record" gorm:"one2many"`
}

type MedicRecord struct {
	PatientDetailsID int `json:"-"`
	SerialNumber string `json:"serial_number" example:"RM/748/2022/0001"`
	Complaint string `json:"complaint" example:"Sakit Perut"`
	Diagnose string `json:"diagnose" example:"Maag"`
	Prescription string `json:"prescription" example:"Entrostop"`
	Doctor string `json:"doctor" example:"Alsyad Ahmad"`
	Facility string `json:"facility" example:"General"`
	DateCheck string `json:"date_check" example:"2022-06-17"`
}

type MedicCheck struct {
	BloodTension int `json:"blood_tension" example:"124"`
	Height int `json:"height" example:"55"`
	Weight int `json:"weight" example:"150"`
	BodyTemperature int `json:"body_temp" example:"34"`
	Nurse string `json:"nurse" example:"Priscilla Halim"`
}
