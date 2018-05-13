package controllers

import (
	"github.com/revel/revel"
	"mcdserver/app/bl"
	"mcdserver/app"
)

type StatisticsController struct {
	*revel.Controller
}

func (controller StatisticsController) Statistics() revel.Result {
	statisticsBl := bl.NewStatisticsBl()

	users, err := statisticsBl.GetUsers()
	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error());
	}

	return controller.RenderJSON(users)
}
