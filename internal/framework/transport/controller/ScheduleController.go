package controller

import (
	"fmt"
	"go-hospital-server/internal/core/entity/request"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/service"
	"go-hospital-server/internal/utils/errors/check"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ScheduleController struct {
	srv *service.ScheduleService
}

func NewScheduleController(srv *service.ScheduleService) *ScheduleController {
	return &ScheduleController{srv: srv}
}

// func (acon ScheduleController) GetScheduleByDate(c echo.Context) error {
// 	date := c.QueryParam("date")
// 	fmt.Println(date)
//
// 	res, err := acon.srv.GetScheduleByDate(date)
// 	if r, ok := check.HTTP(res, err, "Fetch Schedule"); !ok {
// 		return c.JSON(r.Code, r.Result)
// 	}
//
// 	return c.JSON(200, response.MessageData{
// 		Message: "Schedule Fetched",
// 		Data:    res,
// 	})
// }

func (acon ScheduleController) Create(c echo.Context) error {
	var req request.Schedule
	c.Bind(&req)

	if r, ok := check.HTTP(nil, req.Validate(), "Validate"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	err := acon.srv.Create(req)
	if r, ok := check.HTTP(nil, err, "Create Schedule"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(201, response.MessageOnly{
		Message: "Doctor Schedule Created",
	})
}

func (acon ScheduleController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var req request.Schedule
	c.Bind(&req)
	req.SessionID = uint(id)

	if r, ok := check.HTTP(nil, req.Validate(), "Validate"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	err := acon.srv.Update(req)

	if r, ok := check.HTTP(nil, err, "Upadate Doctor Schedule"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageOnly{
		Message: fmt.Sprintf("Doctor Schedule #%d Updated", id),
	})
}

func (acon ScheduleController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := acon.srv.Delete(id)

	if r, ok := check.HTTP(nil, err, "Delete Doctor Schedule"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageOnly{
		Message: fmt.Sprintf("Doctor Schedule #%d Deleted", id),
	})
}
