package mocks

import (
	"go-hospital-server/internal/core/entity/models"
	req "go-hospital-server/internal/core/entity/request"
	res "go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/service"
	"go-hospital-server/internal/utils"
	"go-hospital-server/internal/utils/jwt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var authRepo = &AuthRepository{}
var authService = service.NewAuthService(authRepo)

var token models.Token

func TestLoginSuccess(t *testing.T) {
	pass, _ := utils.HashPassword("password")

	login := req.Login{
		Email: "alsyadahmad@holyhos.id",
		Password: "password",
	}

	user := res.User{
		Email:    "alsyadahmad@holyhos.id",
		Password: pass,
	}

	authRepo.On("FindByEmail", login.Email).Return(user, nil).Once()
	res, err := authService.Login(login)
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}

func TestLoginWrongPW(t *testing.T) {
	pass, _ := utils.HashPassword("password1")

	login := req.Login{
		Email: "alsyadahmad@holyhos.id",
		Password: "password",
	}

	user := res.User{
		Email:    "alsyadahmad@holyhos.id",
		Password: pass,
	}

	authRepo.On("FindByEmail", login.Email).Return(user, nil).Once()
	res, err := authService.Login(login)
	assert.NotEmpty(t, res)
	assert.Error(t, err)
}

func TestRegister(t *testing.T) {
	data := req.UserRequest{
		Email:             "alsyadahmad@holyhos.id",
		Password:          "password",
		FullName:          "Alsyad Ahmad",
		Gender:            "Male",
		RoleID:            2,
		MedicalFacilityID: 1,
	}

	authRepo.On("Register", mock.Anything).Return(nil).Once()
	err := authService.Register(data)
	assert.NoError(t, err)
}

func TestProfile(t *testing.T) {
	data := res.User{
		Code:							 "DCR0001",
		Email:             "alsyadahmad@holyhos.id",
		Password:          "password",
		FullName:          "Alsyad Ahmad",
		Gender:            "Male",
	}

	authRepo.On("FindByCode", mock.Anything).Return(data, nil).Once()
	res, err := authService.MyProfile("DCR0001")
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}

func TestFindEmail(t *testing.T) {
	data := req.FindEmail{
		Email:             "alsyadahmad@holyhos.id",
	}

	resp := res.User{
		Email: "alsyadahmad@holyhos.id",
	}

	authRepo.On("FindByEmail", mock.Anything).Return(resp, nil).Once()
	res, err := authService.FindEmail(data)
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}

func TestChangePW(t *testing.T) {
	data := req.ChangePassword{
		Password: "password",
	}

	authRepo.On("ChangePassword", mock.Anything).Return(nil).Once()
	err := authService.ChangePassword(data)
	assert.NoError(t, err)
}

func TestTokenJWT(t *testing.T) {
	authRepo.On("CreateToken", mock.Anything).Return(mock.Anything, nil).Once()
	authRepo.On("SaveToken", mock.Anything).Return(nil).Once()
	res, err := authService.CreateToken("DCR001", "Doctor", jwt.ACCESS)
	token = res
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}

func TestRefreshToken(t *testing.T) {
	authRepo.On("UpdateToken", token, token).Return(nil).Once()
	res, err := authService.RefreshToken(token)
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
}

func TestRefreshNoToken(t *testing.T) {
	authRepo.On("UpdateToken", token, token).Return(nil).Once()
	token.AccessToken = ""
	res, err := authService.RefreshToken(token)
	assert.Empty(t, res)
	assert.Error(t, err)
}

func TestRevokeToken(t *testing.T) {
	authRepo.On("RevokeToken", mock.Anything).Return(nil).Once()
	err := authService.RevokeToken(token.AccessToken)
	assert.NoError(t, err)
}
