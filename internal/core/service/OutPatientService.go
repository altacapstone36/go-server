package service

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/request"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/repository"
	"go-hospital-server/internal/utils"
	"go-hospital-server/internal/utils/errors/check"
)

type OutPatientService struct {
	repo repository.OutPatientRepository
}

func NewOutPatientService(repo repository.OutPatientRepository) *OutPatientService {
	return &OutPatientService{repo: repo}
}

func (srv OutPatientService) NewMedicRecord(req request.AdminMedicRecord) (err error) {
	var mr models.MedicRecord

	mr, _ = utils.TypeConverter[models.MedicRecord](req)
	mr.MedicalSession, _ = utils.TypeConverter[models.MedicalSession](req)
	mr.Status = 0

	err  = srv.repo.NewMedicalRecord(mr)

	err = check.Record(nil, err)
	return
}

func (srv OutPatientService) Process(req interface{}) (err error) {
	var mr models.MedicRecord

	mr, _ = utils.TypeConverter[models.MedicRecord](req)
	if mr.Diagnose != "" && mr.Prescription != "" {
		mr.Status = 1
	}

	err  = srv.repo.Proceed(mr)

	err = check.Record(nil, err)
	return
}

func (srv OutPatientService) ListPatient(id float64, facility, role string) (res []response.OutPatientResponse, err error) {
	if role == "doctor" {
		res, err  = srv.repo.ListForDoctor(int(id))
	} else if role == "nurse" {
		res, err  = srv.repo.ListForNurse(facility)
	}
	err = check.Record(res, err)
	return
}

func (srv OutPatientService) FilterByDate(start, end string) (res []response.OutPatientResponse, err error) {
	res, err  = srv.repo.FilterByDate(start, end)
	err = check.Record(res, err)
	return
}

func (srv OutPatientService) Report() (res []response.OutPatientReportResponse, err error) {
	res, err  = srv.repo.Report()
	err = check.Record(res, err)
	return
}
