package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-hospital-server/internal/core/service"
	"go-hospital-server/internal/framework/repository"
	"go-hospital-server/internal/framework/routes"
	"go-hospital-server/internal/framework/transport/controller"
	"go-hospital-server/internal/utils/mocktesting"
	"go-hospital-server/internal/utils/mocktesting/test_case"
	"go-hospital-server/internal/utils/validators"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreatePatient(t *testing.T) {
	for _, v := range test_case.NewCreatePatientTest() {
		t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
			
			repo := repository.NewPatientRepository(gdb)
			serv := service.NewPatientService(repo)
			ctrl := controller.NewPatientController(serv)

			e := echo.New()
			api := e.Group("/api")
			validators.NewValidator(gdb)
			routes.NewPatientRoutes(api, ctrl)

			body, _ := json.Marshal(v.Data)
			v.ExpectQuery(sqlmock)

			req := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()


			ctx := e.NewContext(req, rec)
			ctx.SetPath(v.Path)
			err = ctrl.CreatePatient(ctx)

			assert.NoError(t, err)
			assert.Equal(t, v.Code, rec.Code)
		})
	}
}

func TestUpdatePatient(t *testing.T) {
	for _, v := range test_case.NewUpdatePatientTest() {
		t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
		
			repo := repository.NewPatientRepository(gdb)
			serv := service.NewPatientService(repo)
			ctrl := controller.NewPatientController(serv)

			e := echo.New()
			api := e.Group("/api")
			validators.NewValidator(gdb)
			routes.NewPatientRoutes(api, ctrl)

			body, _ := json.Marshal(v.Data)
			v.ExpectQuery(sqlmock)

			req := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()


			ctx := e.NewContext(req, rec)
			ctx.SetPath(v.Path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")
			err = ctrl.UpdatePatient(ctx)

			assert.NoError(t, err)
			assert.Equal(t, v.Code, rec.Code)
		})
	}
}

func TestDeletePatient(t *testing.T) {
	for _, v := range test_case.NewDeletePatientTest() {
		t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
		
			repo := repository.NewPatientRepository(gdb)
			serv := service.NewPatientService(repo)
			ctrl := controller.NewPatientController(serv)

			e := echo.New()
			api := e.Group("/api")
			validators.NewValidator(gdb)
			routes.NewPatientRoutes(api, ctrl)

			body, _ := json.Marshal(v.Data)
			v.ExpectQuery(sqlmock)

			req := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()


			ctx := e.NewContext(req, rec)
			ctx.SetPath(v.Path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")
			err = ctrl.DeletePatient(ctx)

			assert.NoError(t, err)
			assert.Equal(t, v.Code, rec.Code)
		})
	}
}

func TestFindAllPatient(t *testing.T) {
	for _, v := range test_case.NewFindAllPatientTest() {
		t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
		
			repo := repository.NewPatientRepository(gdb)
			serv := service.NewPatientService(repo)
			ctrl := controller.NewPatientController(serv)

			e := echo.New()
			api := e.Group("/api")
			validators.NewValidator(gdb)
			routes.NewPatientRoutes(api, ctrl)

			body, _ := json.Marshal(v.Data)
			v.ExpectQuery(sqlmock)

			req := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", adminToken))
			rec := httptest.NewRecorder()


			ctx := e.NewContext(req, rec)
			ctx.SetPath(v.Path)
			err = ctrl.GetAllPatient(ctx)

			assert.NoError(t, err)
			assert.Equal(t, v.Code, rec.Code)
		})
	}
}

func TestFindPatientByID(t *testing.T) {
	for _, v := range test_case.NewFindPatientByIDTest() {
		t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
		
			repo := repository.NewPatientRepository(gdb)
			serv := service.NewPatientService(repo)
			ctrl := controller.NewPatientController(serv)

			e := echo.New()
			api := e.Group("/api")
			validators.NewValidator(gdb)
			routes.NewPatientRoutes(api, ctrl)

			body, _ := json.Marshal(v.Data)
			v.ExpectQuery(sqlmock)

			req := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", adminToken))
			rec := httptest.NewRecorder()


			ctx := e.NewContext(req, rec)
			ctx.SetPath(v.Path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")
			err = ctrl.GetPatientByID(ctx)

			assert.NoError(t, err)
			assert.Equal(t, v.Code, rec.Code)
		})
	}
}
