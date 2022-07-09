package controller

import (
	"fmt"
	"go-hospital-server/internal/core/entity/request"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/service"
	"go-hospital-server/internal/utils"
	"go-hospital-server/internal/utils/errors/check"
	"go-hospital-server/internal/utils/jwt"
	"strconv"

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
// @Success 200 {object} response.MessageData{data=[]response.OutPatient} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /outpatient [get]
func (acon OutPatientController) GetAllOutPatient(c echo.Context) error {
	token, err := jwt.GetToken(c, jwt.ACCESS)
	role, _ := jwt.GetTokenData(token, "role", jwt.ACCESS)
	code, _ := jwt.GetTokenData(token, "code", jwt.ACCESS)
	var res []response.OutPatient

	date_start := c.QueryParam("date_start")
	date_end := c.QueryParam("date_end")

	if date_start != "" || date_end != "" {
		res, err = acon.srv.FindByDate(code.(string), date_start, date_end)
	} else {
		res, err = acon.srv.ListPatient(code.(string), role.(string))
	}

	if r, ok := check.HTTP(res, err, "Fetch OutPatient"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageData{
		Message: "OutPatient Fetched",
		Data:    res,
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
// @Success 201 {object} response.MessageOnly{} success
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

	return c.JSON(201, response.MessageOnly{
		Message: "Medic Record Created",
	})
}

// godoc
// @Summary Process Doctor
// @Description Process Medic Record
// @Tags OutPatient
// @Security ApiKey
// @Accept json
// @Produce json
// @Param body_doctor body request.DoctorMedicRecord{} false "Process Medic Record by Doctor"
// @Param body_nurse body request.NurseMedicRecord{} false "Process Medic Record by Nurse"
// @Success 200 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /outpatient/:id/process [post]
func (acon OutPatientController) Process(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	token, err := jwt.GetToken(c, jwt.ACCESS)
	role, _ := jwt.GetTokenData(token, "role", jwt.ACCESS)
	var req any
	var process func(int, any) error
	c.Bind(&req)

	if role.(string) == "doctor" {
		r, _ := utils.TypeConverter[request.DoctorMedicRecord](req)
		r.ID = id
		err = r.Validate()
		process = func(id int, t any) error {
			return acon.srv.ProcessDoctor(id, t)
		}
	} else if role.(string) == "nurse" {
		r, _ := utils.TypeConverter[request.NurseMedicRecord](req)
		r.ID = id
		err = r.Validate()
		process = func(id int, t any) error {
			return acon.srv.ProcessNurse(id, t)
		}
	}

	if r, ok := check.HTTP(nil, err, "Validate"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	err = process(id, req)
	if r, ok := check.HTTP(nil, err, "Submit Medic Record"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageOnly{
		Message: "Medic Record Submitted",
	})
}

// godoc
// @Summary Process Doctor
// @Description Process Medic Record
// @Tags OutPatient
// @Security ApiKey
// @Accept json
// @Produce json
// @Param body body request.AssignNurse{} true "Assign Nurse to Medical Check"
// @Success 201 {object} response.MessageOnly{} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /outpatient/:id/assign_nurse [post]
func (acon OutPatientController) AssignNurse(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var req request.AssignNurse
	c.Bind(&req)
	req.ID = id

	if r, ok := check.HTTP(nil, req.Validate(), "Validate"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	err := acon.srv.AssignNurse(id, req)
	if r, ok := check.HTTP(nil, err, "Assign Nurse"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(201, response.MessageOnly{
		Message: fmt.Sprintf("Nurse %s assigned to Medic Record #%d", req.Code, id),
	})
}

// godoc
// @Summary Process Doctor
// @Description Process Medic Record
// @Tags OutPatient
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "Patient ID"
// @Success 200 {object} response.MessageData{data=[]response.OutPatient} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /outpatient/:id [get]
func (acon OutPatientController) FindByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	res, err := acon.srv.FindByID(id)
	if r, ok := check.HTTP(nil, err, "Fetch Medic Record"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageData{
		Message: "Medic Record Fetched",
		Data:    res,
	})
}

// godoc
// @Summary Report Log
// @Description Show Report of Submitted data by Medical Staff
// @Tags OutPatient
// @Security ApiKey
// @Accept json
// @Produce json
// @Success 200 {object} response.MessageData{data=response.OutPatientReportLog} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /outpatient/log [get]
func (acon OutPatientController) ReportLog(c echo.Context) error {
	token, err := jwt.GetToken(c, jwt.ACCESS)
	role, _ := jwt.GetTokenData(token, "role", jwt.ACCESS)
	code, _ := jwt.GetTokenData(token, "code", jwt.ACCESS)

	res, err := acon.srv.ReportLog(code.(string), role.(string))
	if r, ok := check.HTTP(nil, err, "Fetch Report Data"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageData{
		Message: "Report Fetched",
		Data:    res,
	})
}

// godoc
// @Summary Report Log
// @Description Show Report of Submitted data by All Medical Staff
// @Tags OutPatient
// @Security ApiKey
// @Accept json
// @Produce json
// @Success 200 {object} response.MessageData{data=response.OutPatientReportLog} success
// @Failure 417 {object} response.Error{} error
// @Failure 500 {object} response.Error{} error
// @Router /outpatient/report [get]
func (acon OutPatientController) Report(c echo.Context) error {

	res, err := acon.srv.Report()
	if r, ok := check.HTTP(nil, err, "Fetch Report Data"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageData{
		Message: "Report Fetched",
		Data:    res,
	})
}
