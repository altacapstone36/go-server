package test_case

import (
	"go-hospital-server/internal/utils"
	mdt "go-hospital-server/internal/utils/mocktesting/data"
	mqry "go-hospital-server/internal/utils/mocktesting/query"
	"net/http"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewLoginTestCase() []expected {
	return []expected{
		{
			Title:  "No Email",
			Method: http.MethodPost,
			Path:		"/api/login",
			Code:   400,
			Data:   nil,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				_mock.ExpectQuery(mqry.FindUserByEmail()).WithArgs(nil).
					WillReturnRows(_mock.NewRows(nil))
			},
		},
		{
			Title:  "No Record",
			Method: http.MethodPost, Path:		"/api/login",
			Code:   404,
			Data:   mdt.Login,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				_mock.ExpectQuery(mqry.FindUserByEmail()).WithArgs(mdt.Login.Email).
					WillReturnRows(_mock.NewRows(nil))
			},
		},
		{
			Title:  "Wrong Password",
			Method: http.MethodPost,
			Path:		"/api/login",
			Code:   417,
			Data:   mdt.Login,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				row := _mock.NewRows([]string{"email", "password"}).AddRow("alsyadahmad@holyhos.co.id", "")
				_mock.ExpectQuery(mqry.FindUserByEmail()).WithArgs(mdt.Login.Email).
					WillReturnRows(row)
			},
		},
		{
			Title:  "Login Admin",
			Method: http.MethodPost,
			Path:		"/api/login",
			Code:   200,
			Data:   mdt.Login,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				pw, _ := utils.HashPassword("password")
				row := _mock.NewRows([]string{"email", "code", "password", "role", "facility"}).
					AddRow("alsyadahmad@holyhos.co.id", "ADM00001", pw, "Admin", "General")
				_mock.ExpectQuery(mqry.FindUserByEmail()).WithArgs(mdt.Login.Email).
					WillReturnRows(row)
			},
		},
		{
			Title:  "Login Doctor",
			Method: http.MethodPost,
			Path:		"/api/login",
			Code:   200,
			Data:   mdt.Login,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				pw, _ := utils.HashPassword("password")
				row := _mock.NewRows([]string{"email", "code", "password", "role", "facility"}).
					AddRow("alsyadahmad@holyhos.co.id", "DCR00001", pw, "Doctor", "General")
				_mock.ExpectQuery(mqry.FindUserByEmail()).WithArgs(mdt.Login.Email).
					WillReturnRows(row)
			},
		},
		{
			Title:  "Login Nurse",
			Method: http.MethodPost,
			Path:		"/api/login",
			Code:   200,
			Data:   mdt.Login,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				pw, _ := utils.HashPassword("password")
				row := _mock.NewRows([]string{"email", "code", "password", "role", "facility"}).
					AddRow("alsyadahmad@holyhos.co.id", "NRS00001", pw, "Nurse", "General")
				_mock.ExpectQuery(mqry.FindUserByEmail()).WithArgs(mdt.Login.Email).
					WillReturnRows(row)
			},
		},
	}
}

func NewRegisterTestCase() []expected {
	dt := mdt.Register
	return []expected{
		{
			Title:  "Failed to Register",
			Method: http.MethodPost,
			Path:   "/api/register",
			Code:   400,
			Data:   nil, ExpectQuery: func(_mock sqlmock.Sqlmock) {},
		},
		{
			Title:  "Register Success",
			Method: http.MethodPost,
			Path:   "/api/register",
			Code:   201,
			Data:   dt,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				count0 := _mock.NewRows([]string{"COUNT(*)"}).AddRow(0)
				count1 := _mock.NewRows([]string{"COUNT(*)"}).AddRow(1)
				code := _mock.NewRows([]string{"code"}).AddRow("DCR")

				// Validation
				_mock.ExpectQuery(mqry.Duplicate("users", "email", dt.Email, 0)).
					WillReturnRows(count0)
				_mock.ExpectQuery(mqry.Duplicate("users", "full_name", dt.FullName, 0)).
					WillReturnRows(count0)
				_mock.ExpectQuery(mqry.RoleLimit(dt.MedicalFacilityID)).
					WithArgs(mdt.Register.RoleID).
					WillReturnRows(count0)
				_mock.ExpectQuery(mqry.AvailableFacility()).
					WithArgs(dt.MedicalFacilityID).
					WillReturnRows(count1)

				// Register
				any := sqlmock.AnyArg()
				_mock.ExpectBegin()
				_mock.ExpectQuery(mqry.UserRole()).
					WithArgs(dt.RoleID).
					WillReturnRows(count0)
				_mock.ExpectQuery(mqry.RoleCode()).
					WithArgs(dt.RoleID).
					WillReturnRows(code)
				_mock.ExpectExec(mqry.NewUser()).
					WithArgs(any, dt.FullName, dt.Gender, dt.Email, any, any, dt.RoleID, dt.MedicalFacilityID, any).
					WillReturnResult(sqlmock.NewResult(1, 1))
				_mock.ExpectCommit()
			},
		},
	}
}

func NewProfileTest() []expected {
	dt := mdt.User
	return []expected{
		{
			Title:  "Success Get Profile",
			Method: http.MethodGet,
			Path:   "/api/profile",
			Code:   200,
			Data:   dt,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				row := _mock.NewRows([]string{"code", "full_name"}).AddRow("DCR0001", "Alsyad Ahmad")
				_mock.ExpectQuery(mqry.FindUserByCode()).WithArgs(dt.Code).
					WillReturnRows(row)
			},
		},
		{
			Title:  "Failed Get Profile",
			Method: http.MethodGet,
			Path:   "/api/profile",
			Code:   404,
			Data:   dt,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				_mock.ExpectQuery(mqry.FindUserByCode()).WithArgs(dt.Code).
					WillReturnRows(_mock.NewRows(nil))
			},
		},
	}
}

func NewFindEmaileTest() []expected {
	dt := mdt.User
	return []expected{
		{
			Title:  "Success Find Email",
			Method: http.MethodPost,
			Path:   "/api/find_email",
			Code:   200,
			Data:   dt,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				row := _mock.NewRows([]string{"code", "full_name"}).AddRow("DCR00001", "Alsyad Ahmad")
				_mock.ExpectQuery(mqry.FindUserByEmail()).WithArgs(dt.Email).
					WillReturnRows(row)
			},
		},
		{
			Title:  "Failed Find Email",
			Method: http.MethodPost,
			Path:   "/api/find_email",
			Code:   404,
			Data:   dt,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				_mock.ExpectQuery(mqry.FindUserByEmail()).WithArgs(dt.Email).
					WillReturnRows(_mock.NewRows(nil))
			},
		},
	}
}

func NewForgotPWTest() []expected {
	any := sqlmock.AnyArg()
	dt := mdt.ChangePassword
	return []expected{
		{
			Title:  "Success Forgot Password",
			Method: http.MethodPost,
			Path:   "/api/forgot_password?token=",
			Code:   200,
			Data:   dt,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				_mock.ExpectBegin()
				_mock.ExpectExec(mqry.ChangePassword()).WithArgs(dt.Code, any, dt.Code).
					WillReturnResult(sqlmock.NewResult(1, 1))
				_mock.ExpectCommit()
			},
		},
		{
			Title:  "Failed Forgot Password",
			Method: http.MethodPost,
			Path:   "/api/forgot_password?token=",
			Code:   400,
			Data:   nil,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				_mock.ExpectBegin()
				_mock.ExpectExec(mqry.ChangePassword()).WithArgs(dt.Code, any, dt.Code).
					WillReturnResult(sqlmock.NewResult(1, 1))
				_mock.ExpectCommit()
			},
		},
	}
}
