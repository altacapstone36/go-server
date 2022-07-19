package controller

import (
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/service"
	"go-hospital-server/internal/utils/errors/check"

	"github.com/labstack/echo/v4"
)

type SessionController struct {
	srv *service.SessionService
}

func NewSessionController(srv *service.SessionService) *SessionController {
	return &SessionController{srv}
}

func (acon SessionController) FindAllSession(c echo.Context) error {
	res, err := acon.srv.FindAllSession()

	if r, ok := check.HTTP(res, err, "Fetch Session"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageData{
		Message: "Session Fetched",
		Data:    res,
	})
}
