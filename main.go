package main

import (
	"A-brick/app"
	"A-brick/controller"
	"A-brick/middlewarefuncs"
	model "A-brick/model/basemodel"
	"A-brick/model/usermodel"
	"fmt"
	"os"
	"path/filepath"
)

var port = ":8088"
var d = 0  // 显示方式
var b = 40 // 背景色
var f = 31 // 前景色
func main() {
	fmt.Println()
	fmt.Println("------------------------------------")
	fmt.Println("             A birck                ")
	fmt.Printf(" %c[%d;%d;%dm           server-go               %c[0m \n", 0x1B, d, b, f, 0x1B)
	fmt.Println("           port", port, "              ")
	fmt.Println(" 	  -version-0.01-            ")
	fmt.Println("------------------------------------")

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println("path", dir)

	// 注册模型
	model.PoolAdd("login", &usermodel.LoginModel{})
	// 模型服务运行
	model.PoolRun()

	// 注册静态文件服务
	app.Static["/static"] = "static"

	app.Initmiddeware() // 中间件初始化
	// 注册中间件
	// index/Index index=indexcontroller Index=Index
	app.SetMiddlewareFunc("index/Index", middlewarefuncs.Index_middleware)
	app.SetMiddlewareFunc("login/Createuser", middlewarefuncs.Index_middleware)

	// 控制器/路由注册
	app.AutoRouter(&controller.IndexController{})
	app.AutoRouter(&controller.LoginController{})

	app.ServerRun_Http(os.Args[1] + port)
}
