package service

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/request"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/repository"
	"go-hospital-server/internal/utils"
	"go-hospital-server/internal/utils/errors/check"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (srv UserService) Create(req request.UserRequest) (err error) {
	m, _ := utils.TypeConverter[models.User](req)
	err = srv.repo.Create(m)
	err = check.Record(nil, err)
	return
}

func (srv UserService) Update(id int, req request.UserRequest) (err error) {
	m, _ := utils.TypeConverter[models.User](req)
	m.ID = uint(id)
	err = srv.repo.Update(m)
	err = check.Record(nil, err)
	return
}

func (srv UserService) Delete(id int) (err error) {
	err = srv.repo.Delete(id)
	err = check.Record(nil, err)
	return
}

func (srv UserService) FindAll() (user []response.User, err error) {
	user, err = srv.repo.FindAll()
	err = check.Record(user, err)
	return
}

func (srv UserService) FindByID(id int) (user response.User, err error) {
	user, err = srv.repo.FindByID(id)
	err = check.Record(user, err)
	return
}
