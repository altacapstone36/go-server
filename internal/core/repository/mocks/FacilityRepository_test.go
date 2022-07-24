package mocks

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/request"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

var faciRepo = &FacilityRepository{}
var faciService = service.NewFacilityService(faciRepo)

func TestCreateFacility(t *testing.T) {
	r := request.Facility{
		Name: "Apotek",
	}

	data := models.MedicalFacility{Name: "Apotek"}

	faciRepo.On("Create", data).Return(nil).Once()
	err := faciService.Create(r)
	assert.NoError(t, err)
}

func TestUpdateFacility(t *testing.T) {
	r := request.Facility{
		ID: 1,
		Name: "Apotek",
	}

	data := models.MedicalFacility{ID: 1, Name: "Apotek"}

	faciRepo.On("Update", data).Return(nil).Once()
	err := faciService.Update(r)
	assert.NoError(t, err)
}

func TestDeleteFacility(t *testing.T) {
	faciRepo.On("Delete", 1).Return(nil).Once()
	err := faciService.Delete(1)
	assert.NoError(t, err)
}

func TestFindByIDFacility(t *testing.T) {
	resp := response.FacilityDetails{
		ID:       1,
	}

	faciRepo.On("FindByID", 1).Return(resp, nil).Once()
	res, err := faciService.FindByID(1)
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}

func TestFindAllFacility(t *testing.T) {
	resp := []response.Facility{
		{
			ID:       1,
		},
	}


	faciRepo.On("FindAll").Return(resp, nil).Once()
	res, err := faciService.FindAll()
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}
