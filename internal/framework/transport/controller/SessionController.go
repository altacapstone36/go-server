package controller

import (
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/service"
	"go-hospital-server/internal/utils/errors/check"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SessionController struct {
	srv *service.SessionService
}

func NewSessionController(srv *service.SessionService) *SessionController {
	return &SessionController{srv}
}

func (acon SessionController) FindAll(c echo.Context) error {
	res, err := acon.srv.FindAll()

	if r, ok := check.HTTP(res, err, "Fetch Session"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageData{
		Message: "Session Fetched",
		Data:    res,
	})
}

func (acon SessionController) FindByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	date := c.QueryParam("date")
	res, err := acon.srv.FindByDateID(id, date)

	if r, ok := check.HTTP(res, err, "Fetch Session"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageData{
		Message: "Session Fetched",
		Data:    res,
	})
}
