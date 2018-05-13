package bl

import (
	"mcdserver/app/dao"
	"mcdserver/app/models"
)

type StatisticsBl struct {
}

func NewStatisticsBl () StatisticsBl {
	return StatisticsBl{}
}

func (statisticsBl StatisticsBl) getUsersFromDB() ([]models.User, error) {
	userDao := new(dao.UserDao)

	users, err := userDao.GetAllUsers()

	return users, err
}

func (statisticsBl StatisticsBl) GetUsers() ([]models.User, error) {
	users, err := statisticsBl.getUsersFromDB()
	if err != nil {
		return nil, err
	}

	return users, nil
}