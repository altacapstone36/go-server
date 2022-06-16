package service

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/request"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/repository"
	"go-hospital-server/internal/utils/errors/check"
)

type OutPatientService struct {
	repo repository.OutPatientRepository
}

func NewOutPatientService(repo repository.OutPatientRepository) *OutPatientService {
	return &OutPatientService{repo: repo}
}

func (srv OutPatientService) NewMedicRecord(req request.AdminMedicRecord) (err error) {
	var ms models.MedicalSession

	ms.MedicRecord.PatientID = req.PatientID
	ms.MedicRecord.Complaint = req.Complaint
	ms.MedicRecord.MedicalStaffID = req.MedicalStaffID
	ms.SessionID = req.SessionID
	ms.DateCheck = req.DateCheck
	ms.MedicalFacilityID = req.MedicalFacilityID

	err  = srv.repo.NewMedicalRecord(ms)

	err = check.Record(nil, err)
	return
}

func (srv OutPatientService) DoctorProcess(req request.DoctorMedicRequest) (err error) {
	var mr models.MedicRecord

	mr.ID = req.ID
	mr.Diagnose = req.Diagnose
	mr.Prescription = req.Prescription

	err  = srv.repo.Proceed(mr)

	err = check.Record(nil, err)
	return
}

func (srv OutPatientService) NurseProcess(req request.NurseMedicRequest) (err error) {
	var mr models.MedicRecord

	mr.ID = req.ID
	mr.BloodTension = req.BloodTension
	mr.BodyTemperature = req.BodyTemperature
	mr.Weight = req.Weight
	mr.Height = req.Height

	err  = srv.repo.Proceed(mr)

	err = check.Record(nil, err)
	return
}

func (srv OutPatientService) ListPatient(id float64, role string) (res []response.OutPatientResponse, err error) {
	res, err  = srv.repo.ListAvailable(int(id), role)
	err = check.Record(res, err)
	return
}

func (srv OutPatientService) FilterByDate() (res []response.OutPatientResponse, err error) {
	// res, err  = srv.repo.ListAvailable()
	err = check.Record(res, err)
	return
}
