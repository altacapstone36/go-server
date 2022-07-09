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

type UserController struct {
	srv *service.UserService
}

func NewUserController(srv *service.UserService) *UserController {
	return &UserController{srv: srv}
}

// godoc
// @Summary GetAllUser
// @Description Fetch All User Data
// @Tags User
// @Security ApiKey
// @Accept json
// @Produce json
// @Param role_id query string false "User Filter By Role ID"
// @Param facility_id query string false "User Filter By Facility ID"
// @Param session_id query string false "User Filter By Session ID"
// @Success 200 {object} response.MessageData{data=[]response.User} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /user [get]
func (acon UserController) GetAllUser(c echo.Context) error {
	var res []response.User
	var err error
	role_id, _ := strconv.Atoi(c.QueryParam("role_id"))
	facility_id, _ := strconv.Atoi(c.QueryParam("facility_id"))
	session_id, _ := strconv.Atoi(c.QueryParam("session_id"))

	if role_id == 0 && facility_id == 0{
		res, err = acon.srv.FindAll()
	} else {
		res, err = acon.srv.FindByRoleFacility(role_id, facility_id, session_id)
	}

	if r, ok := check.HTTP(res, err, "Fetch User"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageData{
		Message: "User Fetched",
		Data:    res,
	})
}

// godoc
// @Summary GetUserByID
// @Description Fetch User Data By ID
// @Tags User
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} response.MessageData{data=response.User} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /user/:id [get]
func (acon UserController) GetUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := acon.srv.FindByID(id)

	if r, ok := check.HTTP(res, err, "Fetch User"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageData{
		Message: "User Fetched",
		Data:    res,
	})
}

// godoc
// @Summary CreateUser
// @Description Create New User Data
// @Tags User
// @Security ApiKey
// @Accept json
// @Produce json
// @Param body body request.UserRequest true "user data"
// @Success 201 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /user [post]
func (acon UserController) Create(c echo.Context) error {
	var req request.UserRequest
	c.Bind(&req)

	if r, ok := check.HTTP(nil, req.Validate(), "Validate"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	err := acon.srv.Create(req)
	if r, ok := check.HTTP(nil, err, "Create User"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(201, response.MessageOnly{
		Message: "User Created",
	})
}

// godoc
// @Summary UpdateUser
// @Description Update User Data By ID
// @Tags User
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Param body body request.UserRequest true "user data"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /user/:id/update [put]
func (acon UserController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var req request.UserRequest
	c.Bind(&req)
	req.ID = id

	if r, ok := check.HTTP(nil, req.Validate(), "Validate"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	err := acon.srv.Update(req)

	if r, ok := check.HTTP(nil, err, "Update User"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageOnly{
		Message: fmt.Sprintf("User #%d Updated", id),
	})
}

// godoc
// @Summary UpdateUser
// @Description Delete User Data By ID
// @Tags User
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /user/:id/delete [delete]
func (acon UserController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := acon.srv.Delete(id)

	if r, ok := check.HTTP(nil, err, "Delete User"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageOnly{
		Message: fmt.Sprintf("User #%d Deleted", id),
	})
}
