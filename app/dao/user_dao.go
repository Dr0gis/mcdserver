package dao

import (
	"fmt"
	"mcdserver/app/models"
	"database/sql"
	"errors"
	"mcdserver/app"
)

type UserDao struct {
}

type userModelDB struct {
	id int
	email string
	password string
	name sql.NullString
	surname sql.NullString
}

func (userDao UserDao) GetAllUsers() ([]models.User, error) {
	sqlQuery := fmt.Sprintf("SELECT * FROM user")

	rows, err := selectQueryToDataBase(sqlQuery)
	if err != nil {
		return nil, err
	}

	users := make([]*userModelDB, 0)
	for rows.Next() {
		user := new(userModelDB)
		err := rows.Scan(&user.id, &user.email, &user.password, &user.name, &user.surname)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	usersModel := make([]models.User, 0)
	for _, user := range users {
		userModel := models.NewUser(user.email, user.password, user.name.String, user.surname.String)
		usersModel = append(usersModel, userModel)
	}

	return usersModel, nil
}

func (userDao UserDao) GetUserByEmail(email string) (models.User, error) {
	sqlQuery := fmt.Sprintf("SELECT * FROM user WHERE user.email = '%s'", email)

	rows, err := selectQueryToDataBase(sqlQuery)
	if err != nil {
		return models.User{}, err
	}

	user := new(userModelDB)
	countRows := 0
	for rows.Next() {
		countRows++
		err := rows.Scan(&user.id, &user.email, &user.password, &user.name, &user.surname)
		if err != nil {
			return models.User{}, err
		}
	}

	if countRows == 0 {
		return models.User{}, errors.New("user not found")
	}

	userModel := models.NewUser(user.email, user.password, user.name.String, user.surname.String)
	return userModel, nil
}

func (userDao UserDao) InsertUser(email string, password string) error {
	sqlQuery := fmt.Sprintf("INSERT INTO user (email, password) VALUES ('%s', '%s')", email, password)

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