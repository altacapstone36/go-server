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

func (srv OutPatientService) ProcessDoctor(id int, req interface{}) (err error) {
	var mr models.MedicRecord
	mr.ID = uint(id)

	mr, _ = utils.TypeConverter[models.MedicRecord](req)
	mr.Status = 1

	err  = srv.repo.ProceedDoctor(mr)

	err = check.Record(nil, err)
	return
}

func (srv OutPatientService) ProcessNurse(id int, req interface{}) (err error) {
	var mr models.MedicCheck
	mr.ID = uint(id)

	mr, _ = utils.TypeConverter[models.MedicCheck](req)
	mr.Status = 1

	err  = srv.repo.ProceedNurse(mr)

	err = check.Record(nil, err)
	return
}

func (srv OutPatientService) ListPatient(id float64, role string) (res []response.OutPatientResponse, err error) {
	if role == "doctor" {
		res, err  = srv.repo.DoctorFindAll(int(id))
	} else if role == "nurse" {
		res, err  = srv.repo.NurseFindAll(int(id))
	}
	err = check.Record(res, err)
	return
}

func (srv OutPatientService) FindByDate(id float64, start, end string) (res []response.OutPatientResponse, err error) {
	res, err  = srv.repo.FindByDate(start, end)
	err = check.Record(res, err)
	return
}

func (srv OutPatientService) FindByID(id int) (res []response.OutPatientResponse, err error) {
	res, err  = srv.repo.FindByID(id)
	err = check.Record(res, err)
	return
}

func (srv OutPatientService) Report() (res []response.OutPatientReportResponse, err error) {
	res, err  = srv.repo.Report()
	err = check.Record(res, err)
	return
}

func (srv OutPatientService) ReportLog(id float64, role string) (res []response.OutPatientReportLogResponse, err error) {
	res, err  = srv.repo.ReportLog(int(id), role)
	err = check.Record(res, err)
	return
}

func (srv OutPatientService) AssignNurse(id int, mc request.AssignNurseRequest) (err error) {
	mcheck, _ := utils.TypeConverter[models.MedicCheck](mc)
	mcheck.ID = uint(id)
	err = srv.repo.AssignNurse(mcheck)
	return
}
