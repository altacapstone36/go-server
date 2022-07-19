package repository

type Repository struct {
	Auth       AuthRepository
	Patient    PatientRepository
	OutPatient OutPatientRepository
	User       UserRepository
	Facility   FacilityRepository
	Schedule   ScheduleRepository
	Session    SessionRepository
}
