package bl

import (
	"mcdserver/app/dao"
	"mcdserver/app/models"
)

type StatisticsBl struct {
	userEmail string
}

func NewStatisticsBl() StatisticsBl {
	return StatisticsBl{}
}

func NewStatisticsForMovementBl(userEmail string) StatisticsBl {
	return StatisticsBl{userEmail: userEmail}
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

func (statisticsBl StatisticsBl) getDroneMovementsForUserFromDB(iduser int) ([]models.DroneMovement, error) {
	droneMovementDao := new(dao.DroneMovementDao)

	droneMovements, err := droneMovementDao.GetDroneMovementsForUser(iduser)

	return droneMovements, err
}

func (statisticsBl StatisticsBl) GetDroneMovementsForUser() ([]models.DroneMovement, error) {
	userBl := NewUserBl(statisticsBl.userEmail, "")
	user, err := userBl.GetInfo()
	if err != nil {
		return nil, err
	}

	droneMovements, err := statisticsBl.getDroneMovementsForUserFromDB(user.Id)
	if err != nil {
		return nil, err
	}

	zonePointBl := NewZonePointEmptyBl()
	zonePoints, err := zonePointBl.GetZonePoints()
	if err != nil {
		return nil, err
	}

	offenceDroneMovement := make([]models.DroneMovement, 0)
	for _, droneMovement := range droneMovements {
		for _, zonePoint := range zonePoints {
			if droneMovement.X == zonePoint.X && droneMovement.Y == zonePoint.Y {
				if zonePoint.Forbidden && droneMovement.Height >= zonePoint.MinHeight && droneMovement.Height <= zonePoint.MaxHeight {
					offenceDroneMovement = append(offenceDroneMovement, droneMovement)
				}
				continue;
			}
		}
	}

	return offenceDroneMovement, nil
}

func (statisticsBl StatisticsBl) getDroneMovementsFromDB() ([]models.DroneMovement, error) {
	droneMovementDao := new(dao.DroneMovementDao)

	droneMovements, err := droneMovementDao.GetAllDroneMovements()

	return droneMovements, err
}

func (statisticsBl StatisticsBl) getUserByIdDroneActivation(idDroneActivation int) (models.User, error) {
	userDao := dao.UserDao{}

	user, err := userDao.GetUserByIdDroneActivation(idDroneActivation)

	return user, err
}

func (statisticsBl StatisticsBl) GetEvents() ([]models.Event, error) {
	droneMovements, err := statisticsBl.getDroneMovementsFromDB()
	if err != nil {
		return nil, err
	}

	zonePointBl := NewZonePointEmptyBl()
	zonePoints, err := zonePointBl.GetZonePoints()
	if err != nil {
		return nil, err
	}

	events := make([]models.Event, 0)
	for _, droneMovement := range droneMovements {
		for _, zonePoint := range zonePoints {
			if droneMovement.X == zonePoint.X && droneMovement.Y == zonePoint.Y {
				if zonePoint.Forbidden && droneMovement.Height >= zonePoint.MinHeight && droneMovement.Height <= zonePoint.MaxHeight {
					user, err := statisticsBl.getUserByIdDroneActivation(droneMovement.IdDroneActivation)
					if err != nil {
						return nil, err
					}
					events = append(events, models.NewEvent(user.Email, droneMovement.Id, droneMovement.IdDroneActivation, droneMovement.Datetime, droneMovement.X, droneMovement.Y, droneMovement.Height))
				}
				continue;
			}
		}
	}

	return events, nil
}