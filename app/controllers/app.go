package controllers

import (
	"github.com/revel/revel"
	"mcdserver/app"
	"fmt"
	"log"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

type Stuff struct {
	Foo string ` json:"foo" xml:"foo" `
	Bar int ` json:"bar" xml:"bar" `
}

func (c App) Test() revel.Result {
	data := make(map[string] interface{})
	data["error"] = nil
	stuff := Stuff{Foo: "qwe", Bar: 999}
	data["stuff"] = stuff
	data["test"] = "value"

	return c.RenderJSON(data)
}

type User struct {
	id int
	name string
}

func (c App) Home() revel.Result {
	var jsonData map[string] interface {}
	c.Params.BindJSON(&jsonData)

	sql := "SELECT * FROM my_table"
	rows, err := app.DB.Query(sql)

	if (err != nil) {
		fmt.Printf("error db")
	}

	users := make([]*User, 0)
	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.id, &user.name)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	jsonData["users"] = users

	return c.RenderJSON(jsonData)
}
