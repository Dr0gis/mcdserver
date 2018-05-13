package dao

import (
	"database/sql"
	"mcdserver/app"
)

func selectQueryToDataBase(sqlQuery string) (*sql.Rows, error) {
	rows, err := app.DB.Query(sqlQuery)
	return rows, err
}

func insertQueryToDataBase(sqlQuery string) (sql.Result, error) {
	result, err := app.DB.Exec(sqlQuery)
	return result, err
}