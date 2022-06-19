package controller

import (
	"go-hospital-server/internal/core/entity/request"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/service"
	"go-hospital-server/internal/utils"
	"go-hospital-server/internal/utils/errors/check"
	"go-hospital-server/internal/utils/jwt"

	"github.com/labstack/echo/v4"
)

type OutPatientController struct {
	srv *service.OutPatientService
}

func NewOutPatientController(srv *service.OutPatientService) *OutPatientController {
	return &OutPatientController{srv}
}

// godoc
// @Summary GetAllPatient
// @Description Fetch All Patient with New Medic Record Data
// @Tags OutPatient
// @Security ApiKey
// @Accept json
// @Produce json
// @Param date_start query string false "Date Filter Range Start (2022-02-23)"
// @Param date_end query string false "Date Filter Range End (2006-02-25)"
// @Success 200 {object} response.MessageData{data=[]response.OutPatientResponse} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /outpatient [get]
func (acon OutPatientController) GetAllOutPatient(c echo.Context) error {
	role, _ := jwt.GetTokenData(c, "role")
	var res []response.OutPatientResponse
	var err error

	date_start := c.QueryParam("date_start")
	date_end := c.QueryParam("date_end")

	if r, ok := check.ParamQuery(date_start, date_end); !ok {
		return c.JSON(r.Code, r.Result)
	}

	if date_start != "" && date_end != "" {
		res, err = acon.srv.FilterByDate(date_start, date_end)
	} else {
		id, _ := jwt.GetTokenData(c, "user_id")
		facility, _ := jwt.GetTokenData(c, "facility")
		res, err = acon.srv.ListPatient(id.(float64), facility.(string), role.(string))
	}

	if r, ok := check.HTTP(res, err, "Fetch OutPatient"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageData{
		Message: "OutPatient Fetched",
		Data: res,
	})
}

// godoc
// @Summary New Medic Record
// @Description Add New Medic Record for Patient
// @Tags OutPatient
// @Security ApiKey
// @Accept json
// @Produce json
// @Param body body request.AdminMedicRecord{} true "New Medic Record"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /outpatient [post]
func (acon OutPatientController) NewMedicRecord(c echo.Context) error {
	var req request.AdminMedicRecord
	c.Bind(&req)

	if r, ok := check.HTTP(nil, req.Validate(), "Validate"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	err := acon.srv.NewMedicRecord(req)
	if r, ok := check.HTTP(nil, err, "Create Medic Record"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageOnly{
		Message: "Medic Record Created",
	})
}

// godoc
// @Summary Process Doctor
// @Description Process Medic Record by Doctor
// @Tags OutPatient
// @Security ApiKey
// @Accept json
// @Produce json
// @Param body body request.DoctorMedicRequest{} true "Process Medic Record by Doctor"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /outpatient/process [post]
func (acon OutPatientController) Process(c echo.Context) error {
	role, _ := jwt.GetTokenData(c, "role")
	var req any
	var err error
	c.Bind(&req)

	if role.(string) == "doctor" {
		r, _ := utils.TypeConverter[request.DoctorMedicRequest](req)
		err = r.Validate()
	} else if role.(string) == "nurse" {
		r, _ := utils.TypeConverter[request.NurseMedicRequest](req)
		err = r.Validate()
	}

	if r, ok := check.HTTP(nil, err, "Validate"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	err = acon.srv.Process(req)
	if r, ok := check.HTTP(nil, err, "Submit Medic Record"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageOnly{
		Message: "Medic Record Submitted",
	})
}
