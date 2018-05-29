package models

import "time"

type DroneMovement struct {
	Id int `json:"id"`
	IdDroneActivation int `json:"idDroneActivation"`
	Datetime time.Time
	X int `json:"x"`
	Y int `json:"y"`
	Height int `json:"height"`
}

func NewDroneMovement(id int, idDroneActivation int, datetime time.Time, x int, y int, height int) DroneMovement {
	return DroneMovement{Id: id, IdDroneActivation: idDroneActivation, Datetime: datetime, X: x, Y: y, Height: height}
}