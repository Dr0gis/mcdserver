package models

type Drone struct {
	Id int `json:"id"`
	Name string   `json:"name"`
	QrCode string `json:"qrcode"`
}

func NewDrone(id int, name string, qrcode string) Drone {
	return Drone{Id: id, Name: name, QrCode: qrcode}
}