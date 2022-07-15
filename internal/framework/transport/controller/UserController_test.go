package controller_test

import (
	"bytes"
	"encoding/json"
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

func TestCreateUser(t *testing.T) {
	for _, v := range test_case.NewRegisterTestCase() {
		t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
			
			repo := repository.NewUserRepository(gdb)
			serv := service.NewUserService(repo)
			ctrl := controller.NewUserController(serv)

			e := echo.New()
			api := e.Group("/api")
			validators.NewValidator(gdb)
			routes.NewUserRoutes(api, ctrl)

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

func TestUpdateUser(t *testing.T) {
	for _, v := range test_case.NewUserUpdateTest() {
		t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
			
			repo := repository.NewUserRepository(gdb)
			serv := service.NewUserService(repo)
			ctrl := controller.NewUserController(serv)

			e := echo.New()
			api := e.Group("/api")
			validators.NewValidator(gdb)
			routes.NewUserRoutes(api, ctrl)

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

func TestDeleteUser(t *testing.T) {
	for _, v := range test_case.NewUserDeleteTest() {
		t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
			
			repo := repository.NewUserRepository(gdb)
			serv := service.NewUserService(repo)
			ctrl := controller.NewUserController(serv)

			e := echo.New()
			api := e.Group("/api")
			validators.NewValidator(gdb)
			routes.NewUserRoutes(api, ctrl)

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

func TestFindAllUser(t *testing.T) {
	for _, v := range test_case.NewUserFindAllTest() {
		t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
			
			repo := repository.NewUserRepository(gdb)
			serv := service.NewUserService(repo)
			ctrl := controller.NewUserController(serv)

			e := echo.New()
			api := e.Group("/api")
			validators.NewValidator(gdb)
			routes.NewUserRoutes(api, ctrl)

			v.ExpectQuery(sqlmock)

			req := httptest.NewRequest(v.Method, "/", nil)
			rec := httptest.NewRecorder()

			ctx := e.NewContext(req, rec)
			ctx.SetPath(v.Path)
			err = ctrl.GetAllUser(ctx)

			assert.NoError(t, err)
			assert.Equal(t, v.Code, rec.Code)
		})
	}
}

func TestFindUserByID(t *testing.T) {
	for _, v := range test_case.NewUserFindByIDTest() {
		t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
			
			repo := repository.NewUserRepository(gdb)
			serv := service.NewUserService(repo)
			ctrl := controller.NewUserController(serv)

			e := echo.New()
			api := e.Group("/api")
			validators.NewValidator(gdb)
			routes.NewUserRoutes(api, ctrl)

			v.ExpectQuery(sqlmock)

			req := httptest.NewRequest(v.Method, "/", nil)
			rec := httptest.NewRecorder()

			ctx := e.NewContext(req, rec)
			ctx.SetPath(v.Path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")
			err = ctrl.GetUserByID(ctx)

			assert.NoError(t, err)
			assert.Equal(t, v.Code, rec.Code)
		})
	}
}
