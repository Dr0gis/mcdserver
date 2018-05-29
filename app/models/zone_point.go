package models

type ZonePoint struct {
	Id int `json:"id"`
	X int   `json:"x"`
	Y int `json:"y"`
	MinHeight int `json:"minHeight"`
	MaxHeight int `json:"maxHeight"`
	Forbidden bool `json:"forbidden"`
}

func NewZonePoint(id int, x int, y int, minHeight int, maxHeight int, forbidden bool) ZonePoint {
	return ZonePoint{Id: id, X: x, Y: y, MinHeight: minHeight, MaxHeight: maxHeight, Forbidden: forbidden}
}
