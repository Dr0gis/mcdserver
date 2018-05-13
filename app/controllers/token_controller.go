package controllers

import (
	"github.com/revel/revel"
	"mcdserver/app/bl"
	"mcdserver/app"
	"github.com/dgrijalva/jwt-go"
	"errors"
)

type TokenController struct {
	*revel.Controller
}

func ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			app.Logs.Print("Unexpected signing method: %v", token.Header["alg"])
			return nil, errors.New("unexpected signing method")
		}

		return app.GetSecretKey(), nil
	})

	if err != nil {
		return false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		app.Logs.Print(claims["email"], claims["exp"])
		return true, nil
	} else {
		return false, nil
	}
}

type authModel struct {
	email string
	password string
}

func (controller TokenController) Auth() revel.Result {
	var jsonData map[string] string
	controller.Params.BindJSON(&jsonData)

	model := new(authModel);
	model.email = jsonData["email"]
	model.password = jsonData["password"]

	userBl := bl.NewUserBl(model.email, model.password)

	tokenString, err := userBl.GetToken()

	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error());
	}

	return controller.RenderJSON(tokenString)
}

type registerModel struct {
	email string
	password string
}

func (controller TokenController) Register() revel.Result {
	var jsonData map[string] string
	controller.Params.BindJSON(&jsonData)

	model := new(registerModel)
	model.email = jsonData["email"]
	model.password = jsonData["password"]

	userbl := bl.NewUserBl(model.email, model.password)

	err := userbl.Registration()

	if err != nil {
		app.Logs.Print(err)
		controller.Response.SetStatus(400)
		return controller.RenderJSON(err.Error());
	}

	return controller.RenderJSON(nil)
}