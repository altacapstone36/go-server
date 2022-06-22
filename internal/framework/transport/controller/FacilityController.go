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

type FacilityController struct {
	srv *service.FacilityService
}

func NewFacilityController(srv *service.FacilityService) *FacilityController {
	return &FacilityController{srv: srv}
}

// godoc
// @Summary GetAllFacility
// @Description Fetch All Facility Data
// @Tags Facility
// @Security ApiKey
// @Accept json
// @Produce json
// @Success 200 {object} response.MessageData{data=[]response.Facility} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /facility [get]
func (acon FacilityController) GetAllFacility(c echo.Context) error {
	res, err := acon.srv.FindAll()

	if r, ok := check.HTTP(res, err, "Fetch Facility"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageData{
		Message: "Facility Fetched",
		Data: res,
	})
}

// godoc
// @Summary GetFacilityByID
// @Description Fetch Facility Data By ID
// @Tags Facility
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} response.MessageData{data=response.FacilityDetails} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /facility/:id [get]
func (acon FacilityController) GetFacilityByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := acon.srv.FindByID(id)

	if r, ok := check.HTTP(res, err, "Fetch Facility"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageData{
		Message: "Facility Fetched",
		Data: res,
	})
}

// godoc
// @Summary CreateFacility
// @Description Create New Facility Data
// @Tags Facility
// @Security ApiKey
// @Accept json
// @Produce json
// @Param body body request.Facility true "user data"
// @Success 201 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /facility [post]
func (acon FacilityController) Create(c echo.Context) error {
	var req request.Facility
	c.Bind(&req)
	
	if r, ok := check.HTTP(nil, req.Validate(), "Validate"); !ok {
		return c.JSON(r.Code, r.Result)
	}


	err := acon.srv.Create(req)
	if r, ok := check.HTTP(nil, err, "Create Facility"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(201, response.MessageOnly{
		Message: "Facility Created",
	})
}

// godoc
// @Summary UpdateFacility
// @Description Update Facility Data By ID
// @Tags Facility
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Param body body request.Facility true "user data"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /facility/:id/update [put]
func (acon FacilityController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var req request.Facility
	c.Bind(&req)

	err := acon.srv.Update(id, req)

	if r, ok := check.HTTP(nil, err, "Update Facility"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageOnly{
		Message: fmt.Sprintf("Facility #%d Updated", id),
	})
}

// godoc
// @Summary UpdateFacility
// @Description Delete Facility Data By ID
// @Tags Facility
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /facility/:id/delete [delete]
func (acon FacilityController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := acon.srv.Delete(id)

	if r, ok := check.HTTP(nil, err, "Delete Facility"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageOnly{
		Message: fmt.Sprintf("Facility #%d Deleted", id),
	})
}
