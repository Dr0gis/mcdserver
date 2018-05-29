package bl

import (
	"mcdserver/app/models"
	"mcdserver/app/dao"
)

type ZonePointBl struct {
	id int
	x int
	y int
	minHeight int
	maxHeight int
	forbidden bool
}

func NewZonePointEmptyBl() ZonePointBl {
	zonePointsBl := ZonePointBl{}
	return zonePointsBl
}

func NewZonePointUpdateBl(id int, minHeight int, maxHeight int, forbidden bool) ZonePointBl {
	zonePointsBl := ZonePointBl{id: id, minHeight: minHeight, maxHeight: maxHeight, forbidden: forbidden}
	return zonePointsBl
}

func (zonePointBl ZonePointBl) getZonePointsFromDB() ([]models.ZonePoint, error) {
	zonePointDao := new(dao.ZonePointDao)

	zonePoints, err := zonePointDao.GetAllZonePoints()

	return zonePoints, err
}

func (zonePointBl ZonePointBl) GetZonePoints() ([]models.ZonePoint, error) {
	zonePoints, err := zonePointBl.getZonePointsFromDB()
	if err != nil {
		return nil, err
	}

	return zonePoints, nil
}

func (zonePointBl ZonePointBl) insertZonePointInDB() error {
	zonePointDao := new(dao.ZonePointDao)

	err := zonePointDao.InsertZonePoint(zonePointBl.x, zonePointBl.y, zonePointBl.maxHeight, zonePointBl.maxHeight, zonePointBl.forbidden)

	return err
}

func (zonePointBl ZonePointBl) deleteAllZonePointsInDB() error {
	zonePointDao := new(dao.ZonePointDao)

	err := zonePointDao.DeleteAllZonePoint()

	return err
}

func (zonePointBl *ZonePointBl) CreateNewAllPoints() error {
	err := zonePointBl.deleteAllZonePointsInDB()
	if err != nil {
		return err
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j ++ {
			zonePointBl.x = i
			zonePointBl.y = j
			zonePointBl.minHeight = 0
			zonePointBl.maxHeight = 0
			zonePointBl.forbidden = false
			err := zonePointBl.insertZonePointInDB()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (zonePointBl ZonePointBl) updateZonePointInDB() error {
	zonePointDao := new(dao.ZonePointDao)

	err := zonePointDao.UpdateZonePoint(zonePointBl.id, zonePointBl.minHeight, zonePointBl.maxHeight, zonePointBl.forbidden)

	return err
}

func (zonePointBl ZonePointBl) UpdateZonePoint() error {
	err := zonePointBl.updateZonePointInDB()
	return err
}