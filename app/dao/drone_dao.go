package dao

import (
	"mcdserver/app/models"
	"fmt"
	"mcdserver/app"
	"errors"
)

type DroneDao struct {
}

type droneModelDB struct {
	id int
	name string
	qrcode string
}

func (droneDao DroneDao) GetAllDrones() ([]models.Drone, error) {
	sqlQuery := fmt.Sprintf("SELECT * FROM dron")

	rows, err := selectQueryToDataBase(sqlQuery)
	if err != nil {
		return nil, err
	}

	drones := make([]*droneModelDB, 0)
	for rows.Next() {
		drone := new(droneModelDB)
		err := rows.Scan(&drone.id, &drone.qrcode, &drone.name)
		if err != nil {
			return nil, err
		}
		drones = append(drones, drone)
	}

	dronesModel := make([]models.Drone, 0)
	for _, drone := range drones {
		droneModel := models.NewDrone(drone.name, drone.qrcode)
		dronesModel = append(dronesModel, droneModel)
	}

	return dronesModel, nil
}

func (droneDao DroneDao) InsertDrone(name string, qrcode string) error {
	sqlQuery := fmt.Sprintf("INSERT INTO dron (name, qrcode) VALUES ('%s', '%s')", name, qrcode)

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