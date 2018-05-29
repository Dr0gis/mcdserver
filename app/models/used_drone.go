package models

import "time"

type UsedDrone struct {
	IdDroneActivation int `json:"idDroneActivation"`
	DroneId int `json:"droneId"`
	Name string `json:"name"`
	Datetime time.Time `json:"datetime"`
}

func NewUsedDrone(IdDroneActivation int, droneId int, name string, datetime time.Time) UsedDrone {
	return UsedDrone{IdDroneActivation: IdDroneActivation, DroneId: droneId, Name: name, Datetime: datetime}
}