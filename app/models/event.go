package models

import "time"

type Event struct {
	UserEmail string `json:"userEmail"`
	Id int `json:"id"`
	IdDroneActivation int `json:"idDroneActivation"`
	Datetime time.Time
	X int `json:"x"`
	Y int `json:"y"`
	Height int `json:"height"`
}

func NewEvent(userEmail string, id int, idDroneActivation int, datetime time.Time, x int, y int, height int) Event {
	return Event{UserEmail: userEmail, Id: id, IdDroneActivation: idDroneActivation, Datetime: datetime, X: x, Y: y, Height: height}
}