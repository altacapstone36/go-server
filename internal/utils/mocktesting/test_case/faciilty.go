package test_case

import (
	"go-hospital-server/internal/utils/mocktesting/data"
	mqry "go-hospital-server/internal/utils/mocktesting/query"
	"net/http"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewCreateFacilityTest() []expected {
	dt := data.Facility
	return []expected {
		{
			Title:  "Create Fail",
			Method: http.MethodPost,
			Path:   "/api/facility",
			Code:   400,
			Data:   nil,
			ExpectQuery: func( sqlmock.Sqlmock) {},
		},
		{
			Title:  "Create Success",
			Method: http.MethodPost,
			Path:   "/api/facility",
			Code:   201,
			Data:   dt,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				count0 := _mock.NewRows([]string{"COUNT(*)"}).AddRow(0)
				_mock.ExpectQuery(mqry.Duplicate("medical_facilities", "name", dt.Name, 0)).
					WillReturnRows(count0)

				_mock.ExpectBegin()
				_mock.ExpectExec(mqry.NewFacility()).
					WithArgs(dt.Name).
					WillReturnResult(sqlmock.NewResult(1, 1))
				_mock.ExpectCommit()
				_mock.ExpectClose()
			},
		},
	}
}

func NewUpdateFacilityTest() []expected {
	dt := data.Facility
	return []expected {
		{
			Title:  "Update Fail",
			Method: http.MethodPut,
			Path:   "/api/facility/:id/update",
			Code:   400,
			Data:   nil,
			ExpectQuery: func(sqlmock.Sqlmock) {},
		},
		{
			Title:  "Update Success",
			Method: http.MethodPut,
			Path:   "/api/facility/:id/update",
			Code:   200,
			Data:   dt,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				count0 := _mock.NewRows([]string{"COUNT(*)"}).AddRow(0)
				_mock.ExpectQuery(mqry.Duplicate("medical_facilities", "name", dt.Name, dt.ID)).
					WillReturnRows(count0)

				_mock.ExpectBegin()
				_mock.ExpectExec(mqry.UpdateFacility()).
					WithArgs(dt.Name, dt.ID).
					WillReturnResult(sqlmock.NewResult(0, 1))
				_mock.ExpectCommit()
				_mock.ExpectClose()
			},
		},
	}
}

func NewDeleteFacilityTest() []expected {
	dt := data.Facility
	return []expected {
		{
			Title:  "Delete Success",
			Method: http.MethodPut,
			Path:   "/api/facility/:id/delete",
			Code:   200,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				_mock.ExpectBegin()
				_mock.ExpectExec(mqry.DeleteFacility()).
					WithArgs(dt.ID).
					WillReturnResult(sqlmock.NewResult(0, 1))
				_mock.ExpectCommit()
				_mock.ExpectClose()
			},
		},
		{
			Title:  "Delete Fail",
			Method: http.MethodPut,
			Path:   "/api/facility/:id/delete",
			Code:   404,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				_mock.ExpectBegin()
				_mock.ExpectExec(mqry.DeleteFacility()).
					WithArgs(dt.ID).
					WillReturnResult(sqlmock.NewResult(0, 0))
				_mock.ExpectCommit()
				_mock.ExpectClose()
			},
		},
	}
}

func NewFindAllFacilityTest() []expected {
	dt := data.Facility
	return []expected {
		{
			Title:  "Find All Success",
			Method: http.MethodGet,
			Path:   "/api/facility",
			Code:   200,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				row := _mock.NewRows([]string{"id", "name"}).AddRow(dt.ID, dt.Name)
				_mock.ExpectQuery(mqry.FindAllFacility()).
					WillReturnRows(row)
			},
		},
		{
			Title:  "Find All Fail",
			Method: http.MethodGet,
			Path:   "/api/facility",
			Code:   404,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				row := _mock.NewRows(nil)
				_mock.ExpectQuery(mqry.FindAllFacility()).
					WillReturnRows(row)
			},
		},
	}
}

func NewFindFacilityByIDTest() []expected {
	dt := data.Facility
	return []expected {
		{
			Title:  "Find By ID Success",
			Method: http.MethodGet,
			Path:   "/api/facility",
			Code:   200,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				row := _mock.NewRows([]string{"id", "name"}).AddRow(dt.ID, dt.Name)
				_mock.ExpectQuery(mqry.FindByIDFacility()).WithArgs(dt.ID).
					WillReturnRows(row)
				_mock.ExpectQuery(mqry.FindFacilityMember()).
					WillReturnRows(row)
			},
		},
		{
			Title:  "Find By ID Fail",
			Method: http.MethodGet,
			Path:   "/api/facility",
			Code:   404,
			ExpectQuery: func(_mock sqlmock.Sqlmock) {
				row := _mock.NewRows(nil)
				_mock.ExpectQuery(mqry.FindByIDFacility()).WithArgs(dt.ID).
					WillReturnRows(row)
			},
		},
	}
}
