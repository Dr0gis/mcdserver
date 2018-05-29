package models

type DroneActivation struct {
	Id int `json:"id"`
	IdDrone int   `json:"idDrone"`
	IdUser int `json:"idUser"`
}

func NewDroneActivation(id int, idDrone int, idUser int) DroneActivation {
	return DroneActivation{Id: id, IdDrone: idDrone, IdUser: idUser}
}

