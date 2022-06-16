package response


type Patient struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Code string `json:"code"`
	ResidentRegistration string `json:"resident_registration"`
	FullName string `json:"full_name" sql:"unique"`
	Gender string `json:"gender"`
	BirthDate string `json:"birthdate"`
}

type PatientDetails struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Code string `json:"code"`
	ResidentRegistration string `json:"resident_registration"`
	FullName string `json:"full_name" sql:"unique"`
	Address string `json:"address"`
	Gender string `json:"gender"`
	BirthDate string `json:"birthdate"`
	BloodType string `json:"blood_type"`
	MedicRecord []MedicRecord `json:"medic_record"`
}

type MedicRecord struct {
	ID uint `gorm:"primarykey:autoIncrement:false"`
	SerialNumber string `json:"serial_number"`
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

