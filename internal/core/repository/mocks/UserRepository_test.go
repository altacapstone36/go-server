package mocks

import (
	"go-hospital-server/internal/core/entity/request"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepo = &UserRepository{}
var userService = service.NewUserService(userRepo)

func TestCreateUser(t *testing.T) {
	r := request.UserRequest{
		FullName: "Alsyad Ahmad",
	}

	userRepo.On("Create", mock.Anything).Return(nil).Once()
	err := userService.Create(r)
	assert.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	r := request.UserRequest{
		ID: 1,
		FullName: "Alsyad Ahmad",
	}

	userRepo.On("Update", mock.Anything).Return(nil).Once()
	err := userService.Update(r)
	assert.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	userRepo.On("Delete", 1).Return(nil).Once()
	err := userService.Delete(1)
	assert.NoError(t, err)
}

func TestFindByIDUser(t *testing.T) {
	resp := response.User{
		ID:       1,
	}

	userRepo.On("FindByID", 1).Return(resp, nil).Once()
	res, err := userService.FindByID(1)
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}

func TestFindAllUser(t *testing.T) {
	resp := []response.User{
		{
			ID:       1,
		},
	}


	userRepo.On("FindAll").Return(resp, nil).Once()
	res, err := userService.FindAll()
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}

func TestFindByRFSUser(t *testing.T) {
	resp := []response.User{
		{
			ID:       1,
		},
	}

	userRepo.On("FindByRFS", 1, 1, 1).Return(resp, nil).Once()
	res, err := userService.FindByRoleFacility(1, 1, 1)
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}
