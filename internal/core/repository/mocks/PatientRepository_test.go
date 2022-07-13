package mocks

import (
	"go-hospital-server/internal/core/entity/request"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var patientRepo = &PatientRepository{}
var patientService = service.NewPatientService(patientRepo)

func TestCreatePatient(t *testing.T) {
	r := request.Patient{
		FullName: "Alsyad Ahmad",
	}

	patientRepo.On("CreatePatient", mock.Anything).Return(nil).Once()
	err := patientService.CreatePatient(r)
	assert.NoError(t, err)
}

func TestUpdatePatient(t *testing.T) {
	r := request.Patient{
		ID: 1,
		FullName: "Alsyad Ahmad",
	}

	patientRepo.On("UpdatePatient", mock.Anything).Return(nil).Once()
	err := patientService.UpdatePatient(r)
	assert.NoError(t, err)
}

func TestDeletePatient(t *testing.T) {
	patientRepo.On("DeletePatientByID", uint(1)).Return(nil).Once()
	err := patientService.DeletePatient(1)
	assert.NoError(t, err)
}

func TestFindByIDPatient(t *testing.T) {
	resp := response.PatientDetails{
		ID:       1,
	}

	patientRepo.On("GetPatientByID", uint(1)).Return(resp, nil).Once()
	res, err := patientService.GetPatientByID(1)
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}

func TestFindAllPatient(t *testing.T) {
	resp := []response.Patient{
		{
			ID:       1,
		},
	}


	patientRepo.On("GetAllPatient").Return(resp, nil).Once()
	res, err := patientService.GetAllPatient()
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}

func TestFindByRFSPatient(t *testing.T) {
	resp := []response.Patient{
		{
			ID:       1,
		},
	}

	patientRepo.On("GetPatientByName", mock.Anything).Return(resp, nil).Once()
	res, err := patientService.GetPatientByName("Ach")
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}
