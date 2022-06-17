package controller

import (
	"go-hospital-server/internal/core/entity/request"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/service"
	"go-hospital-server/internal/utils/errors/check"
	"go-hospital-server/internal/utils/jwt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
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
// @Success 200 {object} response.MessageData{Data=[]response.OutPatientResponse} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /outpatient [get]
func (acon OutPatientController) GetAllOutPatient(c echo.Context) error {
	id, _ := jwt.GetTokenData(c, "user_id")
	role, _ := jwt.GetTokenData(c, "role")
	var res []response.OutPatientResponse

	date_start := c.QueryParam("date_start")
	date_end := c.QueryParam("date_end")

	err := validation.Validate(&date_start, validation.Date("2006-01-02"))
	if r, ok := check.HTTP(nil, err, "Validate"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	err = validation.Validate(&date_end, validation.Date("2006-01-02"))
	if r, ok := check.HTTP(nil, err, "Validate"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	if date_start != "" && date_end != "" {
		res, err = acon.srv.FilterByDate(date_start, date_end)
	} else {
		res, err = acon.srv.ListPatient(id.(float64), role.(string))
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
// @Param body body request.AdminMedicRecord{} true "new medic record"
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
// @Param body body request.DoctorMedicRequest{} true "process medic record by doctor"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /outpatient/doctor [post]
func (acon OutPatientController) DoctorProcess(c echo.Context) error {
	var req request.DoctorMedicRequest
	c.Bind(&req)

	if r, ok := check.HTTP(nil, req.Validate(), "Validate"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	err := acon.srv.DoctorProcess(req)
	if r, ok := check.HTTP(nil, err, "Submit Medic Record"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageOnly{
		Message: "Medic Record Submitted",
	})
}

// godoc
// @Summary Process Nurse
// @Description Process Medic Record by Doctor
// @Tags OutPatient
// @Security ApiKey
// @Accept json
// @Produce json
// @Param body body request.NurseMedicRequest{} true "process medic record by nurse"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /outpatient/nurse [post]
func (acon OutPatientController) NurseProcess(c echo.Context) error {
	var req request.NurseMedicRequest
	c.Bind(&req)

	if r, ok := check.HTTP(nil, req.Validate(), "Validate"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	err := acon.srv.NurseProcess(req)
	if r, ok := check.HTTP(nil, err, "Submit Medic Record"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageOnly{
		Message: "Medic Record Submitted",
	})
}
