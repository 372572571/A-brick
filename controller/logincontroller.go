package controller

import (
	"A-brick/app"
	"A-brick/model"
)

type LoginController struct {
	app.App
}

func (c *LoginController) Login() {
	m := model.ModelPool["login"]
	m.Go("login", nil)
}
