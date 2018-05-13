package dao

import (
	"fmt"
	"errors"
	"mcdserver/app/models"
	"mcdserver/app"
)

type AdminDao struct {
}

type adminModelDB struct {
	id int
	email string
	login string
	password string
}

func (adminDao AdminDao) GetAdminByEmailOrLogin(emailOrLogin string) (models.Admin, error) {
	sqlQuery := fmt.Sprintf("SELECT * FROM administrator WHERE administrator.email = '%s' OR administrator.login = '%s'", emailOrLogin, emailOrLogin)

	rows, err := selectQueryToDataBase(sqlQuery)
	if err != nil {
		return models.Admin{}, err
	}

	admin := new(adminModelDB)
	countRows := 0
	for rows.Next() {
		countRows++
		err := rows.Scan(&admin.id, &admin.email, &admin.login, &admin.password)
		if err != nil {
			return models.Admin{}, err
		}
	}

	if countRows == 0 {
		return models.Admin{}, errors.New("user not found")
	}

	adminModel := models.NewAdmin(admin.email, admin.login, admin.password)
	return adminModel, nil
}

func (adminDao AdminDao) InsertAdmin(email string, login string, password string) error {
	sqlQuery := fmt.Sprintf("INSERT INTO administrator (email, login, password) VALUES ('%s', '%s', '%s')", email, login, password)

	result, err := insertQueryToDataBase(sqlQuery)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil
	}

	app.Logs.Print(fmt.Sprintf("Rows affected : %d", rowsAffected))
	if rowsAffected == 0 {
		return errors.New("0 rows affected")
	}

	return nil
}