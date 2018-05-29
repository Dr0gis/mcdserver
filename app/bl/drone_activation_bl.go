package bl

import (
	"mcdserver/app/models"
	"mcdserver/app/dao"
)

type DroneActivationBl struct {
	id int
	idDrone int
	idUser int
}

func NewDroneActivationBl(id int, idDrone int, idUser int) DroneActivationBl {
	droneActivationBl := DroneActivationBl{id: id, idDrone: idDrone, idUser: idUser}
	return droneActivationBl
}

func (droneActivationBl DroneActivationBl) getDroneActivationsFromDB() ([]models.DroneActivation, error) {
	droneActivationDao := new(dao.DroneActivationDao)

	droneActivations, err := droneActivationDao.GetDroneActivationsForUser(droneActivationBl.idUser)

	return droneActivations, err
}

func (droneActivationBl DroneActivationBl) getDroneFromDB(idDrone int) (models.Drone, error) {
	droneDao := new(dao.DroneDao)

	drone, err := droneDao.GetDroneById(idDrone)

	return drone, err
}

func (droneActivationBl DroneActivationBl) getDroneMovementsFromDB(idDroneActivation int) ([]models.DroneMovement, error) {
	droneMovementDao := new(dao.DroneMovementDao)

	droneMovements, err := droneMovementDao.GetDroneMovementsForActivation(idDroneActivation)

	return droneMovements, err
}

func (droneActivationBl DroneActivationBl) GetUsedDronesForUser() ([]models.UsedDrone, error) {
	droneActivations, err := droneActivationBl.getDroneActivationsFromDB()
	if err != nil {
		return nil, err
	}

	usedDrones := make([]models.UsedDrone, 0)
	for _, droneActivation := range droneActivations {

		drone, err := droneActivationBl.getDroneFromDB(droneActivation.IdDrone)
		if err != nil {
			return nil, err
		}

		droneMovements, err := droneActivationBl.getDroneMovementsFromDB(droneActivation.Id)
		if err != nil {
			return nil, err
		}

		usedDrone := models.NewUsedDrone(droneActivation.Id, drone.Id, drone.Name, droneMovements[0].Datetime)
		usedDrones = append(usedDrones, usedDrone)
	}

	return usedDrones, nil
}

func (droneActivationBl DroneActivationBl) insertDroneActivationInDB(idDrone int, idUser int) (int, error) {
	droneActivationDao := new(dao.DroneActivationDao)

	id, err := droneActivationDao.InsertDroneActivation(idDrone, idUser)

	return id, err
}

func (droneActivationBl DroneActivationBl) ActivateDrone() (int, error) {
	id, err := droneActivationBl.insertDroneActivationInDB(droneActivationBl.idDrone, droneActivationBl.idUser)

	return id, err
}