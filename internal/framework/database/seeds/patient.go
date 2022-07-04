package seeds

import (
	"go-hospital-server/internal/core/entity/models"

	"gorm.io/gorm"
)

func patientSeeder() Seed {
	seeds := []models.Patient{
		{
			NationalID: "8729301745162748",
			FullName: "Faizur Ramadhan",
			Address: "Sumenep",
			Gender: "Male",
			BirthDate: "2001-02-14",
			BloodType: "A",
			MedicRecord: &[]models.MedicRecord{
				{
					Complaint:       "Sakit Telinga",
					MedicCheck: models.MedicCheck{
						UserCode: "NRS00001",
					},
					MedicalSession: models.MedicalSession{
						MedicalFacilityID: 1,
						UserCode: "DCR00001",
						SessionID:         1,
						DateCheck:         "2022-07-17",
					},
				},
			},
		},
		{
			NationalID: "8761435451636421",
			FullName: "Ach. Novil",
			Address: "Sumenep",
			Gender: "Male",
			BirthDate: "2001-04-21",
			BloodType: "A",
		},
	}
	model := &models.Patient{}

	return Seed{
		models: model,
		run: func(db *gorm.DB) (err error) {
			for _, v := range seeds {
				err = db.Create(&v).Error
			}
			return
		},
	}
}
