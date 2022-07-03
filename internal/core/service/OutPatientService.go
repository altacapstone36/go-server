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
	mr.MedicalSession.UserCode = req.UserCode
	mr.Status = 0

	err = srv.repo.NewMedicalRecord(mr)
	return
}

func (srv OutPatientService) ProcessDoctor(id int, req interface{}) (err error) {
	var mr models.MedicRecord

	mr, _ = utils.TypeConverter[models.MedicRecord](req)
	mr.ID = uint(id)
	mr.Status = 1

	err = srv.repo.ProceedDoctor(mr)
	return
}

func (srv OutPatientService) ProcessNurse(id int, req interface{}) (err error) {
	var mr models.MedicCheck

	mr, _ = utils.TypeConverter[models.MedicCheck](req)
	mr.MedicRecordID = uint(id)
	mr.Status = 1

	err = srv.repo.ProceedNurse(mr)
	return
}

func (srv OutPatientService) ListPatient(code, role string) (res []response.OutPatient, err error) {
	if role == "doctor" {
		res, err = srv.repo.DoctorFindAll(code)
	} else if role == "nurse" {
		res, err = srv.repo.NurseFindAll(code)
	}
	return
}

func (srv OutPatientService) FindByDate(code, start, end string) (res []response.OutPatient, err error) {
	res, err = srv.repo.FindByDate(start, end)
	return
}

func (srv OutPatientService) FindByID(id int) (res response.OutPatientDetails, err error) {
	res, err = srv.repo.FindByID(id)
	return
}

func (srv OutPatientService) Report() (res []response.OutPatientDetails, err error) {
	res, err = srv.repo.Report()
	return
}

func (srv OutPatientService) ReportLog(code, role string) (res []response.OutPatientReportLog, err error) {
	res, err = srv.repo.ReportLog(code, role)
	return
}

func (srv OutPatientService) AssignNurse(id int, mc request.AssignNurse) (err error) {
	mcheck, _ := utils.TypeConverter[models.MedicCheck](mc)
	mcheck.MedicRecordID = uint(id)
	err = srv.repo.AssignNurse(mcheck)
	return
}
