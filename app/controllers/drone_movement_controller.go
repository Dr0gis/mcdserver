package controllers

import (
	"github.com/revel/revel"
	"mcdserver/app"
	"strconv"
	"mcdserver/app/bl"
	"time"
)

type DroneMovementController struct {
	*revel.Controller
}

type droneMovementForUserModel struct {
	token string
	idDroneActivation int
}

func (controller DroneMovementController) DroneMovementForDroneActivation() revel.Result {
	model := droneMovementForUserModel{}
	model.token = controller.Params.Get("token")
	idDroneActivation, err := strconv.ParseInt(controller.Params.Get("idDroneActivation"), 10, 32)
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error());
	}
	model.idDroneActivation = int(idDroneActivation)

	tokenValid, _, _, err := ValidateToken(model.token)
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

	droneMovementBl := bl.NewDroneMovementBl(model.idDroneActivation)
	droneMovements, err := droneMovementBl.GetDroneMovementsForActivation()
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error());
	}

	return controller.RenderJSON(droneMovements)
}

type insertDroneMovementModel struct {
	idDroneActivation int
	datetime time.Time
	x int
	y int
	height int
}

func (controller DroneMovementController) InsertDroneMovement() revel.Result {
	model := new(insertDroneMovementModel)
	idDroneActivation, _ := strconv.ParseInt(controller.Params.Values["idDroneActivation"][0], 10, 32)
	model.idDroneActivation = int(idDroneActivation)
	datetime, _ :=  time.Parse(time.RFC3339, controller.Params.Values["datetime"][0])
	model.datetime = datetime
	x, _ := strconv.ParseInt(controller.Params.Values["x"][0], 10, 32)
	model.x = int(x)
	y, _ := strconv.ParseInt(controller.Params.Values["y"][0], 10, 32)
	model.y = int(y)
	height, _ := strconv.ParseInt(controller.Params.Values["height"][0], 10, 32)
	model.height = int(height)

	droneMovementBl := bl.NewDroneMovementForInsertBl(model.idDroneActivation, model.datetime, model.x, model.y, model.height)
	err := droneMovementBl.InsertDroneMovement()
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error())
	}

	return controller.RenderJSON(nil)
}