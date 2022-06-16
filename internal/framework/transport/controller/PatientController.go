package controller

import (
	"fmt"
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response" 
	"go-hospital-server/internal/core/service"
	"go-hospital-server/internal/utils/errors/check"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)


type PatientController struct {
	srv *service.PatientService
}

func NewPatientController(srv *service.PatientService) *PatientController {
	return &PatientController{srv}
}

// godoc
// @Summary GetAllPatient
// @Description Fetch All Patient Data
// @Tags Patient
// @Security ApiKey
// @Accept json
// @Produce json
// @Success 200 {object} response.MessageData{Data=[]models.Patient} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /patient [get]
func (acon PatientController) GetAllPatient(c echo.Context) error {
	name := c.QueryParam("name")
	var res []response.Patient
	var err error

	if name != "" {
		res, err = acon.srv.GetPatientByName(name)
	} else {
		res, err = acon.srv.GetAllPatient()
	}

	if r, ok := check.HTTP(res, err, "Fetch Patient"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(http.StatusOK, response.MessageData{
		Message: "Patient Fetched",
		Data: res,
	})
}

// godoc
// @Summary GetPatientByID
// @Description Fetch Patient Data By ID
// @Tags Patient
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id  path  string  true "patient id"
// @Success 200 {object} response.MessageData{Data=models.Patient} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /patient/:id [get]
func (acon PatientController) GetPatientByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	res, err := acon.srv.GetPatientByID(uint(id))
	if r, ok := check.HTTP(res, err, "Fetch Patient"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(http.StatusOK, response.MessageData{
		Message: "Patient Fetched",
		Data: res,
	})
}

// godoc
// @Summary CreatePatient
// @Description Fetch All Patient Data
// @Tags Patient
// @Security ApiKey
// @Accept json
// @Produce json
// @Param body  body  models.Patient{}  true "patient details"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /patient [post]
func (acon PatientController) CreatePatient(c echo.Context) error {
	var patient models.Patient

	c.Bind(&patient)
	
	if r, ok := check.HTTP(nil, patient.Validate(), "Validate Patient"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	err := acon.srv.CreatePatient(patient)
	if r, ok := check.HTTP(nil, err, "Created Patient"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(http.StatusCreated, response.MessageOnly{
		Message: "Patient Created",
	})
}

// godoc
// @Summary UpdatePatient
// @Description Fetch All Patient Data
// @Tags Patient
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id  path  int  true "patient id"
// @Param body  body  models.Patient{}  true "patient details"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /patient/:id/update [put]
func (acon PatientController) UpdatePatient(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	var patient models.Patient

	c.Bind(&patient)

	if r, ok := check.HTTP(nil, patient.Validate(), "Validate Patient"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	err = acon.srv.UpdatePatient(uint(id), patient)
	if r, ok := check.HTTP(nil, err, "Update Patient"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(http.StatusOK, response.MessageOnly{
		Message: fmt.Sprintf("Patient #%d Updated", id),
	})
}

// godoc
// @Summary DeletePatient
// @Description Fetch All Patient Data
// @Tags Patient
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id  path  int  true "patient id"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /patient/:id/delete [delete]
func (acon PatientController) DeletePatient(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	err = acon.srv.DeletePatient(uint(id))
	if r, ok := check.HTTP(nil, err, "Update Patient"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(http.StatusOK, response.MessageOnly{
		Message: fmt.Sprintf("Patient #%d Deleted", id),
	})
}
