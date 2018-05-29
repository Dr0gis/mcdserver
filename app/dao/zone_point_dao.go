package dao

import (
	"mcdserver/app/models"
	"fmt"
	"mcdserver/app"
	"errors"
)

type ZonePointDao struct {
}

type zonePointModelDB struct {
	id int
	x int
	y int
	minHeight int
	maxHeight int
	forbidden bool
}

func (zonePointDao ZonePointDao) GetAllZonePoints() ([]models.ZonePoint, error) {
	sqlQuery := fmt.Sprintf("SELECT * FROM zone_points")

	rows, err := selectQueryToDataBase(sqlQuery)
	if err != nil {
		return nil, err
	}

	zonePoints := make([]*zonePointModelDB, 0)
	for rows.Next() {
		zonePoint := new(zonePointModelDB)
		err := rows.Scan(&zonePoint.id, &zonePoint.x, &zonePoint.y, &zonePoint.minHeight, &zonePoint.maxHeight, &zonePoint.forbidden)
		if err != nil {
			return nil, err
		}
		zonePoints = append(zonePoints, zonePoint)
	}

	zonePointsModel := make([]models.ZonePoint, 0)
	for _, zonePoint := range zonePoints {
		zonePointModel := models.NewZonePoint(zonePoint.id, zonePoint.x, zonePoint.y, zonePoint.minHeight, zonePoint.maxHeight, zonePoint.forbidden)
		zonePointsModel = append(zonePointsModel, zonePointModel)
	}

	return zonePointsModel, nil
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (zonePointDao ZonePointDao) InsertZonePoint(x int, y int, minHeight int, maxHeight int, forbidden bool) error {
	sqlQuery := fmt.Sprintf("INSERT INTO zone_points (x, y, min_height, max_height, forbidden) VALUES ('%d', '%d', '%d', '%d', '%d')", x, y, minHeight, maxHeight, bool2int(forbidden))

	result, err := insertQueryToDataBase(sqlQuery)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil
	}

	app.Logs.Print(fmt.Sprintf("Rows affected : %d", rowsAffected))
	if rowsAffected == 0 {
		return errors.New("0 rows affected")
	}

	return nil
}

func (zonePointDao ZonePointDao) DeleteAllZonePoint() error {
	sqlQuery := fmt.Sprintf("DELETE FROM zone_points;")

	result, err := deleteQueryToDataBase(sqlQuery)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil
	}

	app.Logs.Print(fmt.Sprintf("Rows affected : %d", rowsAffected))

	return nil
}

func (zonePointDao ZonePointDao) UpdateZonePoint(id int, minHeight int, maxHeight int, forbidden bool) error {
	sqlQuery := fmt.Sprintf("UPDATE zone_points SET min_height = '%d', max_height = '%d', forbidden = '%d' WHERE idzone_points = '%d'", minHeight, maxHeight, bool2int(forbidden), id)

	result, err := updateQueryToDataBase(sqlQuery)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil
	}

	app.Logs.Print(fmt.Sprintf("Rows affected : %d", rowsAffected))
	if rowsAffected == 0 {
		return errors.New("0 rows affected")
	}

	return nil
}