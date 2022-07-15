package test_case

import (
	"net/http"

	mdt "go-hospital-server/internal/utils/mocktesting/data"
	mqry "go-hospital-server/internal/utils/mocktesting/query"

	"github.com/DATA-DOG/go-sqlmock"
)


func NewUserUpdateTest() []expected {
	dt := mdt.Register
	return []expected{
		{
			Title:  "Failed to Validate",
			Method: http.MethodPut,
			Path:   "/api/user/:id/update",
			Code:   400,
			Data:   nil,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {},
		},
		{
			Title:  "Success Update",
			Method: http.MethodPut,
			Path:   "/api/user/:id/update",
			Code:   200,
			Data:   dt,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				count0 := _mock.NewRows([]string{"COUNT(*)"}).AddRow(0)
				count1 := _mock.NewRows([]string{"COUNT(*)"}).AddRow(1)

				// Validation
				_mock.ExpectQuery(mqry.Duplicate("users", "email", dt.Email, dt.ID)).
					WillReturnRows(count0)
				_mock.ExpectQuery(mqry.Duplicate("users", "full_name", dt.FullName, dt.ID)).
					WillReturnRows(count0)
				_mock.ExpectQuery(mqry.RoleLimit(dt.MedicalFacilityID)).
					WithArgs(mdt.Register.RoleID).
					WillReturnRows(count0)
				_mock.ExpectQuery(mqry.AvailableFacility()).
					WithArgs(dt.MedicalFacilityID).
					WillReturnRows(count1)

				// Update
				any := sqlmock.AnyArg()
				_mock.ExpectBegin()
				_mock.ExpectExec(mqry.UpdateUser()).
					WithArgs(dt.FullName, dt.Gender, dt.Email, any, dt.RoleID, dt.MedicalFacilityID, any).
					WillReturnResult(sqlmock.NewResult(1, 1))
				_mock.ExpectCommit()
			},
		},
	}
}

func NewUserDeleteTest() []expected {
	dt := mdt.Register
	return []expected{
		{
			Title:  "Success Delete",
			Method: http.MethodDelete,
			Path:   "/api/user/:id/delete",
			Code:   200,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				any := sqlmock.AnyArg()
				_mock.ExpectBegin()
				_mock.ExpectExec(mqry.DeleteUser()).
					WithArgs(any, dt.ID).
					WillReturnResult(sqlmock.NewResult(0, 1))
				_mock.ExpectCommit()
			},
		},
		{
			Title:  "Failed Delete",
			Method: http.MethodDelete,
			Path:   "/api/user/:id/delete",
			Code:   404,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				any := sqlmock.AnyArg()
				_mock.ExpectBegin()
				_mock.ExpectExec(mqry.DeleteUser()).
					WithArgs(any, dt.ID).
					WillReturnResult(sqlmock.NewResult(0, 0))
				_mock.ExpectCommit()
			},
		},
	}
}

func NewUserFindAllTest() []expected {
	dt := mdt.User
	return []expected{
		{
			Title:  "Success Find",
			Method: http.MethodGet,
			Path:   "/api/user",
			Code:   200,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				row := _mock.NewRows([]string{"role", "full_name"}).AddRow(dt.Code, dt.FullName)
				_mock.ExpectQuery(mqry.FindAllUser()).
					WillReturnRows(row)
			},
		},
		{
			Title:  "Failed Find",
			Method: http.MethodGet,
			Path:   "/api/user",
			Code:   404,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				_mock.ExpectQuery(mqry.FindAllUser()).
					WillReturnRows(_mock.NewRows(nil))
			},
		},
	}
}

func NewUserFindByIDTest() []expected {
	dt := mdt.User
	return []expected{
		{
			Title:  "Success Find",
			Method: http.MethodGet,
			Path:   "/api/user",
			Code:   200,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				row := _mock.NewRows([]string{"role", "full_name"}).AddRow(dt.Code, dt.FullName)
				_mock.ExpectQuery(mqry.FindUserByID()).
					WillReturnRows(row)
			},
		},
		{
			Title:  "Failed Find",
			Method: http.MethodGet,
			Path:   "/api/user",
			Code:   404,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				_mock.ExpectQuery(mqry.FindUserByID()).
					WillReturnRows(_mock.NewRows(nil))
			},
		},
	}
}
