package dao

import (
	"mcdserver/app/models"
	"fmt"
	"mcdserver/app"
	"errors"
	"time"
	"github.com/revel/revel"
)

type DroneMovementDao struct {
}

type droneMovementModelDB struct {
	id int
	idDroneActivation int
	datetime time.Time
	x int
	y int
	height int
}

func (droneMovementDao DroneMovementDao) GetAllDroneMovements() ([]models.DroneMovement, error) {
	sqlQuery := fmt.Sprintf("SELECT * FROM dron_movement")

	rows, err := selectQueryToDataBase(sqlQuery)
	if err != nil {
		return nil, err
	}

	droneMovements := make([]*droneMovementModelDB, 0)
	for rows.Next() {
		droneMovement := new(droneMovementModelDB)
		err := rows.Scan(&droneMovement.id, &droneMovement.x, &droneMovement.datetime, &droneMovement.idDroneActivation, &droneMovement.y, &droneMovement.height)
		if err != nil {
			return nil, err
		}
		droneMovements = append(droneMovements, droneMovement)
	}

	droneMovementsModel := make([]models.DroneMovement, 0)
	for _, droneMovement := range droneMovements {
		droneMovementModel := models.NewDroneMovement(droneMovement.id, droneMovement.idDroneActivation, droneMovement.datetime, droneMovement.x, droneMovement.y, droneMovement.height)
		droneMovementsModel = append(droneMovementsModel, droneMovementModel)
	}

	return droneMovementsModel, nil
}

func (droneMovementDao DroneMovementDao) GetDroneMovementsForActivation(idDroneActivation int) ([]models.DroneMovement, error) {
	sqlQuery := fmt.Sprintf("SELECT * FROM dron_movement WHERE iddron_activation = '%d' ORDER BY datetime", idDroneActivation)

	rows, err := selectQueryToDataBase(sqlQuery)
	if err != nil {
		return nil, err
	}

	droneMovements := make([]*droneMovementModelDB, 0)
	for rows.Next() {
		droneMovement := new(droneMovementModelDB)
		err := rows.Scan(&droneMovement.id, &droneMovement.x, &droneMovement.datetime, &droneMovement.idDroneActivation, &droneMovement.y, &droneMovement.height)
		if err != nil {
			return nil, err
		}
		droneMovements = append(droneMovements, droneMovement)
	}

	droneMovementsModel := make([]models.DroneMovement, 0)
	for _, droneMovement := range droneMovements {
		droneMovementModel := models.NewDroneMovement(droneMovement.id, droneMovement.idDroneActivation, droneMovement.datetime, droneMovement.x, droneMovement.y, droneMovement.height)
		droneMovementsModel = append(droneMovementsModel, droneMovementModel)
	}

	return droneMovementsModel, nil
}

func (droneMovementDao DroneMovementDao) GetDroneMovementsForUser(iduser int) ([]models.DroneMovement, error) {
	sqlQuery := fmt.Sprintf("SELECT * FROM mcddb.dron_movement WHERE mcddb.dron_movement.iddron_activation IN (SELECT mcddb.dron_activation.iddron_activation FROM mcddb.dron_activation WHERE iduser = %d);", iduser)

	rows, err := selectQueryToDataBase(sqlQuery)
	if err != nil {
		return nil, err
	}

	droneMovements := make([]*droneMovementModelDB, 0)
	for rows.Next() {
		droneMovement := new(droneMovementModelDB)
		err := rows.Scan(&droneMovement.id, &droneMovement.x, &droneMovement.datetime, &droneMovement.idDroneActivation, &droneMovement.y, &droneMovement.height)
		if err != nil {
			return nil, err
		}
		droneMovements = append(droneMovements, droneMovement)
	}

	droneMovementsModel := make([]models.DroneMovement, 0)
	for _, droneMovement := range droneMovements {
		droneMovementModel := models.NewDroneMovement(droneMovement.id, droneMovement.idDroneActivation, droneMovement.datetime, droneMovement.x, droneMovement.y, droneMovement.height)
		droneMovementsModel = append(droneMovementsModel, droneMovementModel)
	}

	return droneMovementsModel, nil
}

func (droneMovementDao DroneMovementDao) InsertDroneMovement(idDroneActivation int, datetime time.Time, x int, y int, height int) error {
	sqlQuery := fmt.Sprintf("INSERT INTO dron_movement (iddron_activation, datetime, x, y, height) VALUES ('%d', '%s', '%d', '%d', '%d')", idDroneActivation, datetime.Format(revel.DateTimeFormat), x, y, height)

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
