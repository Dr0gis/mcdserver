package controllers

import (
	"github.com/revel/revel"
	"mcdserver/app"
	"mcdserver/app/bl"
)

type DroneActivationController struct {
	*revel.Controller
}

type droneActivationForUserModel struct {
	token string
}

func (controller DroneActivationController) UsedDronesForUser() revel.Result {
	model := droneActivationForUserModel{}
	model.token = controller.Params.Get("token")

	tokenValid, _, email, err := ValidateToken(model.token)
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error())
	}
	if !tokenValid {
		app.Logs.Print("token don't valid")
		controller.Response.SetStatus(400)
		return controller.RenderJSON("token don't valid")
	}

	userBl := bl.NewUserBl(email, "")
	user, err := userBl.GetInfo()
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error())
	}

	droneActivationBl := bl.NewDroneActivationBl(-1, -1, user.Id)
	usedDrones, err := droneActivationBl.GetUsedDronesForUser()
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error())
	}

	return controller.RenderJSON(usedDrones)
}

type droneActivated struct {
	Qrcode string `json:"qrcode"`
	IdDroneActivation int `json:"idDroneActivation"`
}

var drones []droneActivated

func (controller DroneActivationController) CheckActivateDrone() revel.Result {
	qrcode := controller.Params.Get("qrcode")
	for _, element := range drones {
		if element.Qrcode == qrcode {
			return controller.RenderJSON(element)
		}
	}

	return controller.RenderJSON(false)
}

type activateDroneModel struct {
	token string
	qrcode string
}

func (controller DroneActivationController) ActivateDrone() revel.Result {
	var jsonData map[string] string
	controller.Params.BindJSON(&jsonData)

	model := new(activateDroneModel)
	model.token = jsonData["token"]
	model.qrcode = jsonData["qrcode"]

	tokenValid, _, email, err := ValidateToken(model.token)
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error());
	}
	if !tokenValid {
		app.Logs.Print("token don't valid")
		controller.Response.SetStatus(400)
		return controller.RenderJSON("token don't valid")
	}

	userBl := bl.NewUserBl(email, "")
	user, err := userBl.GetInfo()
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error())
	}

	droneBl := bl.NewDroneByQrCodeBl(model.qrcode)
	drone, err := droneBl.GetDroneByQrCode()
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error())
	}

	droneActivationBl := bl.NewDroneActivationBl(-1, drone.Id, user.Id)
	idDroneActivation, err := droneActivationBl.ActivateDrone()
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error())
	}

	drones = append(drones, droneActivated{Qrcode: model.qrcode, IdDroneActivation: idDroneActivation})

	return controller.RenderJSON(nil)
}

func (controller DroneActivationController) DeactivateDrone() revel.Result {
	qrcode := controller.Params.Get("qrcode")
	i := -1;
	for index, element := range drones {
		if element.Qrcode == qrcode {
			i = index
			break
		}
	}
	if i != -1 {
		drones = append(drones[:i], drones[i+1:]...)
	}

	return controller.RenderJSON(nil)
}