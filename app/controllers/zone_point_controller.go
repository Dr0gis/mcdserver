package controllers

import (
	"github.com/revel/revel"
	"mcdserver/app"
	"mcdserver/app/bl"
)

type ZonePointController struct {
	*revel.Controller
}

type zonePointsModels struct {
	token string
}

func (controller ZonePointController) ZonePoints() revel.Result {
	model := zonePointsModels{}
	model.token = controller.Params.Get("token")

	tokenValid, isAdmin, _, err := ValidateToken(model.token)
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error());
	}
	if !tokenValid {
		app.Logs.Print("token don't valid")
		controller.Response.SetStatus(400)
		return controller.RenderJSON("token don't valid");
	}
	if !isAdmin {
		app.Logs.Print("user not admin")
		controller.Response.SetStatus(400)
		return controller.RenderJSON("user not admin");
	}

	zonePointBl := bl.NewZonePointEmptyBl()

	zonePoints, err := zonePointBl.GetZonePoints()
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error());
	}

	return controller.RenderJSON(zonePoints)
}

type clearZonePointsModels struct {
	token string
}

func (controller ZonePointController) ClearZonePoints() revel.Result {
	var jsonData map[string] string
	controller.Params.BindJSON(&jsonData)

	model := clearZonePointsModels{}
	model.token = jsonData["token"]

	tokenValid, isAdmin, _, err := ValidateToken(model.token)
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error());
	}
	if !tokenValid {
		app.Logs.Print("token don't valid")
		controller.Response.SetStatus(400)
		return controller.RenderJSON("token don't valid");
	}
	if !isAdmin {
		app.Logs.Print("user not admin")
		controller.Response.SetStatus(400)
		return controller.RenderJSON("user not admin");
	}

	zonePointBl := bl.NewZonePointEmptyBl()

	err = zonePointBl.CreateNewAllPoints()
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error());
	}

	return controller.RenderJSON(nil)
}

type updateZonePointsModels struct {
	token string
	id int
	minHeight int
	maxHeight int
	forbidden bool
}

func (controller ZonePointController) UpdateZonePoints() revel.Result {
	var jsonData map[string] interface{}
	controller.Params.BindJSON(&jsonData)

	model := updateZonePointsModels{}
	model.token = jsonData["token"].(string)
	model.id = int(jsonData["id"].(float64))
	model.minHeight = int(jsonData["minHeight"].(float64))
	model.maxHeight = int(jsonData["maxHeight"].(float64))
	model.forbidden = jsonData["forbidden"].(bool)

	tokenValid, isAdmin, _, err := ValidateToken(model.token)
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error());
	}
	if !tokenValid {
		app.Logs.Print("token don't valid")
		controller.Response.SetStatus(400)
		return controller.RenderJSON("token don't valid");
	}
	if !isAdmin {
		app.Logs.Print("user not admin")
		controller.Response.SetStatus(400)
		return controller.RenderJSON("user not admin");
	}

	zonePointBl := bl.NewZonePointUpdateBl(model.id, model.minHeight, model.maxHeight, model.forbidden)

	err = zonePointBl.UpdateZonePoint()
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error());
	}

	return controller.RenderJSON(nil)
}
