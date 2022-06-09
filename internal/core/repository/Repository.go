package repository

type Repository struct {
	Auth AuthRepository
	Patient PatientRepository
}

