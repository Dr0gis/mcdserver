package bl

import (
	"mcdserver/app/models"
	"mcdserver/app/dao"
	"errors"
	"crypto/sha256"
	"encoding/base64"
)

type DroneBl struct {
	name string
	qrcode string
}

func NewDroneEmptyBl() DroneBl {
	droneBl := DroneBl{}
	return droneBl
}

func NewDroneBl(name string) DroneBl {
	droneBl := DroneBl{name: name}
	return droneBl
}

func (droneBl DroneBl) getDronesFromDB() ([]models.Drone, error) {
	droneDao := new(dao.DroneDao)

	drones, err := droneDao.GetAllDrones()

	return drones, err
}

func (droneBl DroneBl) GetDrones() ([]models.Drone, error) {
	drones, err := droneBl.getDronesFromDB()
	if err != nil {
		return nil, err
	}

	return drones, nil
}

func (droneBl DroneBl) InsertDroneInDB() error {
	droneDao := new(dao.DroneDao)

	err := droneDao.InsertDrone(droneBl.name, droneBl.qrcode)

	return err
}

func (droneBl *DroneBl) createQrCode() {
	hasher := sha256.New()
	hasher.Write([]byte(droneBl.name))
	droneBl.qrcode = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func (droneBl DroneBl) AddNewDrone() error {
	if droneBl.name == "" {
		return errors.New("name must'n empty")
	}

	droneBl.createQrCode()

	err := droneBl.InsertDroneInDB()
	if err != nil {
		return err
	}

	return nil
}