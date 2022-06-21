package request

type Patient struct {
	Code string `json:"code" example:"RM0001"`
	ResidentRegistration string `json:"resident_registration" example:"8729301745162748"`
	FullName string `json:"full_name" sql:"unique" example:"Faizur Ramadhan"`
	Address string `json:"address" example:"Sumenep"`
	Gender string `json:"gender" example:"Male"`
	BirthDate string `json:"birthdate" example:"2001-04-14"`
	BloodType string `json:"blood_type" example:"A"`
}

