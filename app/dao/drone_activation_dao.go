package dao

import (
	"mcdserver/app/models"
	"fmt"
	"mcdserver/app"
	"errors"
)

type DroneActivationDao struct {
}

type droneActivationModelDB struct {
	id int
	idDrone int
	idUser int
}


func (droneActivationDao DroneActivationDao) GetAllDroneActivations() ([]models.DroneActivation, error) {
	sqlQuery := fmt.Sprintf("SELECT * FROM dron_activation")

	rows, err := selectQueryToDataBase(sqlQuery)
	if err != nil {
		return nil, err
	}

	droneActivations := make([]*droneActivationModelDB, 0)
	for rows.Next() {
		droneActivation := new(droneActivationModelDB)
		err := rows.Scan(&droneActivation.id, &droneActivation.idDrone, &droneActivation.idUser)
		if err != nil {
			return nil, err
		}
		droneActivations = append(droneActivations, droneActivation)
	}

	droneActivationsModel := make([]models.DroneActivation, 0)
	for _, droneActivation := range droneActivations {
		droneActivationModel := models.NewDroneActivation(droneActivation.id, droneActivation.idDrone, droneActivation.idUser)
		droneActivationsModel = append(droneActivationsModel, droneActivationModel)
	}

	return droneActivationsModel, nil
}

func (droneActivationDao DroneActivationDao) GetDroneActivationsForUser(idUser int) ([]models.DroneActivation, error) {
	sqlQuery := fmt.Sprintf("SELECT * FROM dron_activation WHERE iduser = '%d'", idUser)

	rows, err := selectQueryToDataBase(sqlQuery)
	if err != nil {
		return nil, err
	}

	droneActivations := make([]*droneActivationModelDB, 0)
	for rows.Next() {
		droneActivation := new(droneActivationModelDB)
		err := rows.Scan(&droneActivation.id, &droneActivation.idDrone, &droneActivation.idUser)
		if err != nil {
			return nil, err
		}
		droneActivations = append(droneActivations, droneActivation)
	}

	droneActivationsModel := make([]models.DroneActivation, 0)
	for _, droneActivation := range droneActivations {
		droneActivationModel := models.NewDroneActivation(droneActivation.id, droneActivation.idDrone, droneActivation.idUser)
		droneActivationsModel = append(droneActivationsModel, droneActivationModel)
	}

	return droneActivationsModel, nil
}

func (droneActivationDao DroneActivationDao) InsertDroneActivation(idDrone int, idUser int) (int, error) {
	sqlQuery := fmt.Sprintf("INSERT INTO dron_activation (iddron, iduser) VALUES ('%d', '%d')", idDrone, idUser)

	result, err := insertQueryToDataBase(sqlQuery)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	app.Logs.Print(fmt.Sprintf("Rows affected : %d", rowsAffected))
	if rowsAffected == 0 {
		return -1, errors.New("0 rows affected")
	}

	return int(id), nil
}