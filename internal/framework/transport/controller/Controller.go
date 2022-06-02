package controller

import "go-hospital-server/internal/core/service"

type Controller struct {
	Auth *AuthController
}

func NewController(srv *service.Service) *Controller {
	return &Controller{
		Auth: NewAuthController(srv.Auth),
	}
}
