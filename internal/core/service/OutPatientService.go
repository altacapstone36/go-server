package service

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/request"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/repository"
	"go-hospital-server/internal/utils"
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
	return
}

func (srv OutPatientService) ProcessDoctor(id int, req interface{}) (err error) {
	var mr models.MedicRecord

	mr, _ = utils.TypeConverter[models.MedicRecord](req)
	mr.ID = uint(id)
	mr.Status = 1

	err  = srv.repo.ProceedDoctor(mr)
	return
}

func (srv OutPatientService) ProcessNurse(id int, req interface{}) (err error) {
	var mr models.MedicCheck

	mr, _ = utils.TypeConverter[models.MedicCheck](req)
	mr.ID = uint(id)
	mr.Status = 1

	err  = srv.repo.ProceedNurse(mr)
	return
}

func (srv OutPatientService) ListPatient(id float64, role string) (res []response.OutPatient, err error) {
	if role == "doctor" {
		res, err  = srv.repo.DoctorFindAll(int(id))
	} else if role == "nurse" {
		res, err  = srv.repo.NurseFindAll(int(id))
	}
	return
}

func (srv OutPatientService) FindByDate(id float64, start, end string) (res []response.OutPatient, err error) {
	res, err  = srv.repo.FindByDate(start, end)
	return
}

func (srv OutPatientService) FindByID(id int) (res []response.OutPatientSimple, err error) {
	res, err  = srv.repo.FindByID(id)
	return
}

func (srv OutPatientService) Report() (res []response.OutPatientReport, err error) {
	res, err  = srv.repo.Report()
	return
}

func (srv OutPatientService) ReportLog(id float64, role string) (res []response.OutPatientReportLog, err error) {
	res, err  = srv.repo.ReportLog(int(id), role)
	return
}

func (srv OutPatientService) AssignNurse(id int, mc request.AssignNurse) (err error) {
	mcheck, _ := utils.TypeConverter[models.MedicCheck](mc)
	mcheck.ID = uint(id)
	err = srv.repo.AssignNurse(mcheck)
	return
}
