package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-hospital-server/internal/core/entity/response"
	"go-hospital-server/internal/core/service"
	"go-hospital-server/internal/framework/repository"
	"go-hospital-server/internal/framework/routes"
	"go-hospital-server/internal/framework/transport/controller"
	"go-hospital-server/internal/framework/transport/middleware"
	"go-hospital-server/internal/utils/mocktesting"
	"go-hospital-server/internal/utils/mocktesting/test_case"
	"go-hospital-server/internal/utils/validators"
	"log"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var adminToken string
var doctorToken string
var nurseToken string
var fpToken string

func TestLogin(t *testing.T) {
	for _, v := range test_case.NewLoginTestCase() { t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
			
			repo := repository.NewRepository(gdb, nil)
			serv := service.NewService(repo)
			ctrl := controller.NewController(serv)

			e := echo.New()
			api := e.Group("/api")
			routes.NewAuthRoutes(api, ctrl.Auth)

			body, _ := json.Marshal(v.Data)
			v.ExpectQuery(sqlmock)

			req := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()


			ctx := e.NewContext(req, rec)
			ctx.SetPath(v.Path)
			err = ctrl.Auth.Login(ctx)

			if strings.Contains(v.Title, "Login") {
				var result response.MessageDataJWT
				if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
						log.Fatalln(err)
				}
				if strings.Contains(v.Title, "Admin") {
					adminToken = result.JWT.AccessToken
				} else if strings.Contains(v.Title, "Doctor") {
					doctorToken = result.JWT.AccessToken
				} else if strings.Contains(v.Title, "Nurse") {
					nurseToken = result.JWT.AccessToken
				}
			}

			assert.Equal(t, v.Code, rec.Code)
		})
	}
}

func TestRegister(t *testing.T) {
	for _, v := range test_case.NewRegisterTestCase() {
		t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
			
			repo := repository.NewAuthRepository(gdb, nil)
			serv := service.NewAuthService(repo)
			ctrl := controller.NewAuthController(serv)

			e := echo.New()
			api := e.Group("/api")
			validators.NewValidator(gdb)
			routes.NewAuthRoutes(api, ctrl)

			body, _ := json.Marshal(v.Data)
			v.ExpectQuery(sqlmock)

			req := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()


			ctx := e.NewContext(req, rec)
			ctx.SetPath(v.Path)
			err = ctrl.Register(ctx)

			assert.NoError(t, err)
			assert.Equal(t, v.Code, rec.Code)
		})
	}
}

func TestProfile(t *testing.T) {
	for _, v := range test_case.NewProfileTest() {
		t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
			
			repo := repository.NewAuthRepository(gdb, nil)
			serv := service.NewAuthService(repo)
			ctrl := controller.NewAuthController(serv)

			e := echo.New()
			api := e.Group("/api")
			api.Use(middleware.JWT)
			validators.NewValidator(gdb)
			routes.NewAuthRoutes(api, ctrl)

			body, _ := json.Marshal(v.Data)
			v.ExpectQuery(sqlmock)

			req := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", doctorToken))
			rec := httptest.NewRecorder()


			ctx := e.NewContext(req, rec)
			ctx.SetPath(v.Path)
			err = ctrl.Profile(ctx)

			assert.NoError(t, err)
			assert.Equal(t, v.Code, rec.Code)
		})
	}
}

func TestFindEmail(t *testing.T) {
	for _, v := range test_case.NewFindEmaileTest() {
		t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
			
			repo := repository.NewAuthRepository(gdb, nil)
			serv := service.NewAuthService(repo)
			ctrl := controller.NewAuthController(serv)

			e := echo.New()
			api := e.Group("/api")
			api.Use(middleware.JWT)
			validators.NewValidator(gdb)
			routes.NewAuthRoutes(api, ctrl)

			body, _ := json.Marshal(v.Data)
			v.ExpectQuery(sqlmock)

			req := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()


			ctx := e.NewContext(req, rec)
			ctx.SetPath(v.Path)
			err = ctrl.FindEmail(ctx)

			if strings.Contains(v.Title, "Success") {
				var result response.MessageDataJWT
				if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
						log.Fatalln(err)
				}

				fpToken = result.JWT.AccessToken
			}

			assert.NoError(t, err)
			assert.Equal(t, v.Code, rec.Code)
		})
	}
}

func TestForgotPW(t *testing.T) {
	for _, v := range test_case.NewForgotPWTest() {
		t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
			
			repo := repository.NewAuthRepository(gdb, nil)
			serv := service.NewAuthService(repo)
			ctrl := controller.NewAuthController(serv)

			e := echo.New()
			api := e.Group("/api")
			api.Use(middleware.JWT)
			validators.NewValidator(gdb)
			routes.NewAuthRoutes(api, ctrl)

			body, _ := json.Marshal(v.Data)
			v.ExpectQuery(sqlmock)

			req := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()


			ctx := e.NewContext(req, rec)
			ctx.SetPath(v.Path)
			ctx.QueryParams().Add("token", fpToken)
			err = ctrl.ForgotPassword(ctx)

			assert.NoError(t, err)
			assert.Equal(t, v.Code, rec.Code)
		})
	}
}

func TestChangePW(t *testing.T) {
	for _, v := range test_case.NewForgotPWTest() {
		t.Run(v.Title, func(t *testing.T) {
			gdb, sqlmock, err := mocktesting.InitGORMSQLMock()
			if err != nil {
				panic(err)
			}
			
			repo := repository.NewAuthRepository(gdb, nil)
			serv := service.NewAuthService(repo)
			ctrl := controller.NewAuthController(serv)

			e := echo.New()
			api := e.Group("/api")
			api.Use(middleware.JWT)
			validators.NewValidator(gdb)
			routes.NewAuthRoutes(api, ctrl)

			body, _ := json.Marshal(v.Data)
			v.ExpectQuery(sqlmock)

			req := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", doctorToken))
			rec := httptest.NewRecorder()


			ctx := e.NewContext(req, rec)
			ctx.SetPath(v.Path)
			err = ctrl.ChangePassword(ctx)

			assert.NoError(t, err)
			assert.Equal(t, v.Code, rec.Code)
		})
	}
}
