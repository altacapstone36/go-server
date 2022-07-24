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

func TestCreateFacility(t *testing.T) {
	for _, v := range test_case.NewCreateFacilityTest() {
		t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
			
			repo := repository.NewFacilityRepository(gdb)
			serv := service.NewFacilityService(repo)
			ctrl := controller.NewFacilityController(serv)

			e := echo.New()
			api := e.Group("/api")
			validators.NewValidator(gdb)
			routes.NewFacilityRoutes(api, ctrl)

			body, _ := json.Marshal(v.Data)
			v.ExpectQuery(sqlmock)

			req := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()


			ctx := e.NewContext(req, rec)
			ctx.SetPath(v.Path)
			err = ctrl.Create(ctx)

			assert.NoError(t, err)
			assert.Equal(t, v.Code, rec.Code)
		})
	}
}

func TestUpdateFacility(t *testing.T) {
	for _, v := range test_case.NewUpdateFacilityTest() {
		t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
			
			repo := repository.NewFacilityRepository(gdb)
			serv := service.NewFacilityService(repo)
			ctrl := controller.NewFacilityController(serv)

			e := echo.New()
			api := e.Group("/api")
			validators.NewValidator(gdb)
			routes.NewFacilityRoutes(api, ctrl)

			body, _ := json.Marshal(v.Data)
			v.ExpectQuery(sqlmock)

			req := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()


			ctx := e.NewContext(req, rec)
			ctx.SetPath(v.Path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")
			err = ctrl.Update(ctx)

			assert.NoError(t, err)
			assert.Equal(t, v.Code, rec.Code)
		})
	}
}

func TestDeleteFacility(t *testing.T) {
	for _, v := range test_case.NewDeleteFacilityTest() {
		t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
			
			repo := repository.NewFacilityRepository(gdb)
			serv := service.NewFacilityService(repo)
			ctrl := controller.NewFacilityController(serv)

			e := echo.New()
			api := e.Group("/api")
			validators.NewValidator(gdb)
			routes.NewFacilityRoutes(api, ctrl)

			body, _ := json.Marshal(v.Data)
			v.ExpectQuery(sqlmock)

			req := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()


			ctx := e.NewContext(req, rec)
			ctx.SetPath(v.Path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")
			err = ctrl.Delete(ctx)

			assert.NoError(t, err)
			assert.Equal(t, v.Code, rec.Code)
		})
	}
}

func TestFindAllFacility(t *testing.T) {
	for _, v := range test_case.NewFindAllFacilityTest() {
		t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
			
			repo := repository.NewFacilityRepository(gdb)
			serv := service.NewFacilityService(repo)
			ctrl := controller.NewFacilityController(serv)

			e := echo.New()
			api := e.Group("/api")
			validators.NewValidator(gdb)
			routes.NewFacilityRoutes(api, ctrl)

			body, _ := json.Marshal(v.Data)
			v.ExpectQuery(sqlmock)

			req := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", adminToken))
			rec := httptest.NewRecorder()


			ctx := e.NewContext(req, rec)
			ctx.SetPath(v.Path)
			err = ctrl.GetAllFacility(ctx)

			assert.NoError(t, err)
			assert.Equal(t, v.Code, rec.Code)
		})
	}
}

func TestFindFacilityByID(t *testing.T) {
	for _, v := range test_case.NewFindFacilityByIDTest() {
		t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
			
			repo := repository.NewFacilityRepository(gdb)
			serv := service.NewFacilityService(repo)
			ctrl := controller.NewFacilityController(serv)

			e := echo.New()
			api := e.Group("/api")
			validators.NewValidator(gdb)
			routes.NewFacilityRoutes(api, ctrl)

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
			err = ctrl.GetFacilityByID(ctx)

			assert.NoError(t, err)
			assert.Equal(t, v.Code, rec.Code)
		})
	}
}
