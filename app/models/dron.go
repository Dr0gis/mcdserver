package models

type Drone struct {
	Name string   `json:"name"`
	QrCode string `json:"qrcode"`
}

func NewDrone(name string, qrcode string) Drone {
	return Drone{Name: name, QrCode: qrcode}
}