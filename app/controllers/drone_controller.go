package controllers

import (
	"github.com/revel/revel"
	"mcdserver/app"
	"mcdserver/app/bl"
)

type DronesController struct {
	*revel.Controller
}

type dronesModel struct {
	token string
}

func (controller DronesController) Drones() revel.Result {
	model := dronesModel{}
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

	dronesBl := bl.NewDroneEmptyBl()

	drones, err := dronesBl.GetDrones()
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error());
	}

	return controller.RenderJSON(drones)
}

type addDroneModel struct {
	token string
	name string
}

func (controller DronesController) AddDrone() revel.Result {
	var jsonData map[string] string
	controller.Params.BindJSON(&jsonData)

	model := new(addDroneModel)
	model.token = jsonData["token"]
	model.name = jsonData["name"]

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

	droneBl := bl.NewDroneBl(model.name)

	err = droneBl.AddNewDrone()
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error());
	}

	return controller.RenderJSON(nil)
}