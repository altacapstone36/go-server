package service

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/request"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/repository"
	"go-hospital-server/internal/utils"
	"go-hospital-server/internal/utils/errors"
	"go-hospital-server/internal/utils/jwt"
)

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (srv AuthService) Login(login request.Login) (res response.User, err error) {
	var checkPassword bool
	res, err = srv.repo.Login(login.Email)

	if err == nil {
		checkPassword = utils.ComparePassword(login.Password, res.Password)
		if !checkPassword {
			err = errors.New(417, "Wrong Password")
		}
	}

	return
}

func (srv AuthService) Register(req request.RegisterRequest) (err error) {
	r, _ := utils.TypeConverter[models.User](req)
	err = srv.repo.Register(r)
	return
}

func (srv AuthService) FindEmail(find request.FindEmail) (err error) {
	err = srv.repo.FindEmail(find.Email)

	if err != nil {
		errors.New(404, "record not found")
	}

	return
}

func (srv AuthService) ChangePassword(change request.ChangePassword) (res response.User, err error) {
	var checkPassword bool
	res, err = srv.repo.ChangePassword(change.Password)

	if err == nil {
		checkPassword = utils.ComparePassword(change.Password, res.Password)
		if !checkPassword {
			err = errors.New(425, "Try another password")
		}
	}
	return
}

func (srv AuthService) Logout(token string) (err error) {
	err = srv.repo.RevokeToken(token)
	return
}

func (srv AuthService) CreateToken(code, facility_id, level string) (t models.Token, err error) {
	t, _ = jwt.CreateToken(code, facility_id, level)
	err = srv.repo.SaveToken(t)
	return
}

func (srv AuthService) RefreshToken(tkn models.Token) (token models.Token, err error) {

	if tkn.AccessToken == "" {
		err = errors.New(401, "Token Not Provided")
		return
	}

	new_token, err := jwt.RefreshToken(tkn)
	if err != nil {
		return
	}
	err = srv.repo.UpdateToken(tkn, new_token)
	token = new_token
	return
}
