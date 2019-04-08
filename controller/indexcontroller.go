package controller

import (
	"A-brick/app"
)

type IndexController struct {
	app.App
}

func (i *IndexController) Index() {
	i.Data["name"] = "A brick"
	i.Data["email"] = "372572571@qq.com"
	i.Echo("view/index.html")
}
