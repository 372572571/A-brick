package controller

import (
	"A-brick/app"
	"A-brick/model"
)

type IndexController struct {
	app.App
}

func (i *IndexController) Index() {
	m := model.ModelPool["login"]
	m.Go("login", nil)
	i.Data["name"] = "A brick"
	i.Data["email"] = "372572571@qq.com"
	i.Echo("view/index.html")
}
