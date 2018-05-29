package bl

import (
	"mcdserver/app/models"
	"mcdserver/app/dao"
	"time"
)

type DroneMovementBl struct {
	idDroneActivation int
	datetime time.Time
	x int
	y int
	height int
}

func NewDroneMovementBl(idDroneActivation int) DroneMovementBl {
	droneMovementBl := DroneMovementBl {idDroneActivation: idDroneActivation}
	return droneMovementBl
}

func NewDroneMovementForInsertBl(idDroneActivation int, datetime time.Time, x int, y int, height int) DroneMovementBl {
	droneMovementBl := DroneMovementBl {idDroneActivation: idDroneActivation, datetime: datetime, x: x, y: y, height: height}
	return droneMovementBl
}

func (droneMovementBl DroneMovementBl) getDroneMovementsFromDB() ([]models.DroneMovement, error) {
	droneMovementDao := new(dao.DroneMovementDao)

	droneMovements, err := droneMovementDao.GetDroneMovementsForActivation(droneMovementBl.idDroneActivation)

	return droneMovements, err
}

func (droneMovementBl DroneMovementBl) GetDroneMovementsForActivation() ([]models.DroneMovement, error) {
	droneMovements, err := droneMovementBl.getDroneMovementsFromDB()
	if err != nil {
		return nil, err
	}

	return droneMovements, nil
}

func (droneMovementBl DroneMovementBl) insertDroneMovementInDB() error {
	droneMovementDao := new(dao.DroneMovementDao)

	err := droneMovementDao.InsertDroneMovement(droneMovementBl.idDroneActivation, droneMovementBl.datetime, droneMovementBl.x, droneMovementBl.y, droneMovementBl.height)

	return err
}

func (droneMovementBl DroneMovementBl) InsertDroneMovement() error {
	err := droneMovementBl.insertDroneMovementInDB()

	return err
}