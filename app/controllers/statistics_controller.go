package controllers

import (
	"github.com/revel/revel"
	"mcdserver/app/bl"
	"mcdserver/app"
)

type StatisticsController struct {
	*revel.Controller
}

type usersListModel struct {
	token string
}

func (controller StatisticsController) UsersList() revel.Result {
	model := usersListModel{}
	model.token = controller.Params.Get("token");

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

	statisticsBl := bl.NewStatisticsBl()

	users, err := statisticsBl.GetUsers()
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error());
	}

	return controller.RenderJSON(users)
}

type statisticsModel struct {
	token string
	emailUser string
}

func (controller StatisticsController) Statistics() revel.Result {
	model := statisticsModel{}
	model.token = controller.Params.Get("token");
	model.emailUser = controller.Params.Get("emailUser")

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

	statisticsBl := bl.NewStatisticsForMovementBl(model.emailUser)

	droneMovements, err := statisticsBl.GetDroneMovementsForUser()
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error());
	}

	return controller.RenderJSON(droneMovements)
}

type eventsModel struct {
	token string
}

func (controller StatisticsController) Events() revel.Result {
	model := eventsModel{}
	model.token = controller.Params.Get("token");

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

	statisticsBl := bl.NewStatisticsBl()

	events, err := statisticsBl.GetEvents()
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error());
	}

	return controller.RenderJSON(events)
}