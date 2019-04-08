package controller

import (
	"A-brick/app"
	"A-brick/app/toolbox/returnjson"
	model "A-brick/model/basemodel"
	"A-brick/model/usermodel"
	"encoding/json"
	"fmt"
	"net/http"
)

// LoginController 登陆控制器
type LoginController struct {
	app.App
}

// Createuser ... 创建随机用户
func (c *LoginController) Createuser() {
	fmt.Println(c.R().Method)
	// OPTIONS 请求处理
	if c.R().Method == "OPTIONS" {
		fmt.Fprintln(c.W(), "")
		return
	}
	if c.R().Method != "POST" {
		http.NotFound(c.W(), c.R()) // 返回404
		return
	}

	m := model.ModelPool["login"]
	res := m.Go("create_rand_user", nil)
	fmt.Println(res.Get(0).(*usermodel.User))
	s, _ := json.Marshal(res.Get(0))
	j := returnjson.NewMassage(1, string(s))
	returnjson.PutJson(c.W(), j)
}
