package data

import "go-hospital-server/internal/core/entity/request"

var Patient request.Patient = request.Patient{
	ID:         1,
	NationalID: "1234567895643728",
	FullName:   "Syarif Ubaidillah",
	Address:    "Batuan",
	Gender:     "Male",
	BirthDate:  "2001-07-02",
	BloodType:  "O",
}
