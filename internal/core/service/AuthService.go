package service

import (
	m "go-hospital-server/internal/core/entity/models"
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
	res, err = srv.repo.FindByEmail(login.Email)

	if err == nil {
		checkPassword = utils.ComparePassword(login.Password, res.Password)
		if !checkPassword {
			err = errors.New(417, "Wrong Password")
		}
	}

	return
}

func (srv AuthService) Register(req request.UserRequest) (err error) {
	r, _ := utils.TypeConverter[m.User](req)
	r.Password, _ = utils.HashPassword(req.Password)
	err = srv.repo.Register(r)
	return
}

func (srv AuthService) MyProfile(code string) (res response.User, err error) {
	res, err = srv.repo.FindByCode(code)
	return
}

func (srv AuthService) FindEmail(find request.FindEmail) (res response.User, err error) {
	res, err = srv.repo.FindByEmail(find.Email)
	return
}

func (srv AuthService) ChangePassword(change request.ChangePassword) (err error) {
	var cp m.User
	cp.Code = change.Code
	cp.Password, _ = utils.HashPassword(change.Password)
	err = srv.repo.ChangePassword(cp)
	return
}

func (srv AuthService) RevokeToken(token string) (err error) {
	err = srv.repo.RevokeToken(token)
	return
}

func (srv AuthService) CreateToken(code, level string, tokenType jwt.Token) (t m.Token, err error) {
	t, _ = jwt.CreateToken(code, level, tokenType)
	err = srv.repo.SaveToken(t)
	return
}

func (srv AuthService) RefreshToken(tkn m.Token) (token m.Token, err error) {

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
