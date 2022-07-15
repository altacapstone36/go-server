package test_case

import "github.com/DATA-DOG/go-sqlmock"

type expected struct {
	Title string
	Method string
	Path string
	Code int
	Data interface{}
	ExpectQuery func(sqlmock.Sqlmock)
}
