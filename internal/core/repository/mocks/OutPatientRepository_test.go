package mocks

import (
	req "go-hospital-server/internal/core/entity/request"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var outpatientRepo = &OutPatientRepository{}
var outpatientService = service.NewOutPatientService(outpatientRepo)

func TestNewRecord(t *testing.T) {
	req_data := req.AdminMedicRecord{
		PatientCode:       "RM0001",
	}

	outpatientRepo.On("NewMedicalRecord", mock.Anything).Return(nil).Once()
	err := outpatientService.NewMedicRecord(req_data)
	assert.NoError(t, err)
}

func TestProcessDoctor(t *testing.T) {
	req_data := req.DoctorMedicRecord{
		ID: 1,
	}

	outpatientRepo.On("ProceedDoctor", mock.Anything).Return(nil).Once()
	err := outpatientService.ProcessDoctor(req_data)
	assert.NoError(t, err)
}

func TestProcessNurse(t *testing.T) {
	req_data := req.NurseMedicRecord{
		ID:              1,
	}

	outpatientRepo.On("ProceedNurse", mock.Anything).Return(nil).Once()
	err := outpatientService.ProcessNurse(req_data)
	assert.NoError(t, err)
}

func TestListAdminOutPatient(t *testing.T) {
	data := []response.OutPatient{
		{
			ID: 1,
		},
	}
	outpatientRepo.On("AdminFindAll").Return(data, nil).Once()
	res, err := outpatientService.ListPatient("ADM001", "admin")
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}

func TestListDoctorOutPatient(t *testing.T) {
	data := []response.OutPatient{
		{
			ID: 1,
		},
	}
	outpatientRepo.On("DoctorFindAll", mock.Anything).Return(data, nil).Once()
	res, err := outpatientService.ListPatient("DCR001", "doctor")
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}

func TestListNurseOutPatient(t *testing.T) {
	data := []response.OutPatient{
		{
			ID: 1,
		},
	}
	outpatientRepo.On("NurseFindAll", mock.Anything).Return(data, nil).Once()
	res, err := outpatientService.ListPatient("NRS001", "nurse")
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}

func TestOutPatientByID(t *testing.T) {
	data := response.OutPatientDetails{
		ID: 1,
	}
	outpatientRepo.On("FindByID", mock.Anything).Return(data, nil).Once()
	res, err := outpatientService.FindByID(1)
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}

func TestOutPatientByDate(t *testing.T) {
	data := []response.OutPatient{
		{
			ID: 1,
		},
	}
	outpatientRepo.On("FindByDate", mock.Anything, mock.Anything, mock.Anything).Return(data, nil).Once()
	res, err := outpatientService.FindByDate("DCR001", "2022-07-01", "2022-07-31")
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}

func TestOutPatientReport(t *testing.T) {
	data := []response.OutPatientDetails{
		{
			ID: 1,
		},
	}
	outpatientRepo.On("Report").Return(data, nil).Once()
	res, err := outpatientService.Report()
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}

func TestOutPatientLog(t *testing.T) {
	data := []response.OutPatientReportLog{
		{
			SerialNumber: "RM/752/2022/0001",
		},
	}
	outpatientRepo.On("ReportLog", mock.Anything, mock.Anything).Return(data, nil).Once()
	res, err := outpatientService.ReportLog("DCR001", "doctor")
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}

func TestOutPatientAssignNurse(t *testing.T) {
	r := req.AssignNurse {
		ID: 1,
		Code: "NRS0001",
	}
	outpatientRepo.On("AssignNurse", mock.Anything).Return(nil).Once()
	err := outpatientService.AssignNurse(1, r)
	assert.NoError(t, err)
}
