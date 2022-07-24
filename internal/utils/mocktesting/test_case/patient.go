package test_case

import (
	"go-hospital-server/internal/utils/mocktesting/data"
	mqry "go-hospital-server/internal/utils/mocktesting/query"
	"net/http"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewCreatePatientTest() []expected {
	dt := data.Patient
	return []expected {
		{
			Title:  "Create Fail",
			Method: http.MethodPost,
			Path:   "/api/patient",
			Code:   400,
			Data:   nil,
			ExpectQuery: func( sqlmock.Sqlmock) {},
		},
		{
			Title:  "Create Success",
			Method: http.MethodPost,
			Path:   "/api/patient",
			Code:   201,
			Data:   dt,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				count0 := _mock.NewRows([]string{"COUNT(*)"}).AddRow(0)
				_mock.ExpectQuery(mqry.Duplicate("patients", "national_id", dt.NationalID, 0)).
					WillReturnRows(count0)
				_mock.ExpectQuery(mqry.Duplicate("patients", "full_name", dt.FullName, 0)).
					WillReturnRows(count0)

				any := sqlmock.AnyArg()
				_mock.ExpectBegin()
				_mock.ExpectExec(mqry.NewPatient()).
					WithArgs(any, dt.NationalID, dt.FullName, dt.Address, dt.Gender, dt.BirthDate, dt.BloodType, any).
					WillReturnResult(sqlmock.NewResult(1, 1))
				_mock.ExpectCommit()
				_mock.ExpectClose()
			},
		},
	}
}

func NewUpdatePatientTest() []expected {
	dt := data.Patient
	return []expected {
		{
			Title:  "Update Fail",
			Method: http.MethodPut,
			Path:   "/api/patient/:id/update",
			Code:   400,
			Data:   nil,
			ExpectQuery: func(sqlmock.Sqlmock) {},
		},
		{
			Title:  "Update Success",
			Method: http.MethodPut,
			Path:   "/api/patient/:id/update",
			Code:   200,
			Data:   dt,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				count0 := _mock.NewRows([]string{"COUNT(*)"}).AddRow(0)
				_mock.ExpectQuery(mqry.Duplicate("patients", "national_id", dt.NationalID, dt.ID)).
					WillReturnRows(count0)
				_mock.ExpectQuery(mqry.Duplicate("patients", "full_name", dt.FullName, dt.ID)).
					WillReturnRows(count0)

				_mock.ExpectBegin()
				_mock.ExpectExec(mqry.UpdatePatient()).
					WithArgs(dt.NationalID, dt.FullName, dt.Address, dt.Gender, dt.BirthDate, dt.BloodType, dt.ID).
					WillReturnResult(sqlmock.NewResult(0, 1))
				_mock.ExpectCommit()
				_mock.ExpectClose()
			},
		},
	}
}

func NewDeletePatientTest() []expected {
	dt := data.Patient
	return []expected {
		{
			Title:  "Delete Success",
			Method: http.MethodPut,
			Path:   "/api/patient/:id/delete",
			Code:   200,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				any := sqlmock.AnyArg()
				_mock.ExpectBegin()
				_mock.ExpectExec(mqry.DeletePatient()).
					WithArgs(any, dt.ID).
					WillReturnResult(sqlmock.NewResult(0, 1))
				_mock.ExpectCommit()
				_mock.ExpectClose()
			},
		},
		{
			Title:  "Delete Fail",
			Method: http.MethodPut,
			Path:   "/api/patient/:id/delete",
			Code:   404,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				any := sqlmock.AnyArg()
				_mock.ExpectBegin()
				_mock.ExpectExec(mqry.DeletePatient()).
					WithArgs(any, dt.ID).
					WillReturnResult(sqlmock.NewResult(0, 0))
				_mock.ExpectCommit()
				_mock.ExpectClose()
			},
		},
	}
}

func NewFindAllPatientTest() []expected {
	dt := data.Patient
	return []expected {
		{
			Title:  "Find All Success",
			Method: http.MethodGet,
			Path:   "/api/patient",
			Code:   200,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				row := _mock.NewRows([]string{"id", "name"}).AddRow(dt.ID, dt.FullName)
				_mock.ExpectQuery(mqry.FindAllPatient()).
					WillReturnRows(row)
			},
		},
		{
			Title:  "Find All Fail",
			Method: http.MethodGet,
			Path:   "/api/patient",
			Code:   404,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				row := _mock.NewRows(nil)
				_mock.ExpectQuery(mqry.FindAllPatient()).
					WillReturnRows(row)
			},
		},
	}
}

func NewFindPatientByIDTest() []expected {
	dt := data.Patient
	return []expected {
		{
			Title:  "Find By ID Success",
			Method: http.MethodGet,
			Path:   "/api/patient",
			Code:   200,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				row := _mock.NewRows([]string{"id", "name"}).AddRow(dt.ID, dt.FullName)
				_mock.ExpectQuery(mqry.FindByIDPatient()).WithArgs(dt.ID).
					WillReturnRows(row)
				_mock.ExpectQuery(mqry.FindPatientMedicRecord()).
					WillReturnRows(row)
			},
		},
		{
			Title:  "Find By ID Fail",
			Method: http.MethodGet,
			Path:   "/api/patient",
			Code:   404,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				row := _mock.NewRows(nil)
				_mock.ExpectQuery(mqry.FindByIDPatient()).WithArgs(dt.ID).
					WillReturnRows(row)
			},
		},
	}
}
