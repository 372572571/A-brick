package main

import(
	"webapp/app"
	"webapp/controller"
	"fmt"
	"webapp/middlewarefuncs"
)

var port=":8082"
var d=0 // 显示方式
var b=40 // 背景色
var f=31 // 前景色
func main(){
	fmt.Println()
	fmt.Println("------------------------------------")
	fmt.Println("             A birck                ")
	fmt.Printf(" %c[%d;%d;%dm           server-go               %c[0m \n", 0x1B, d, b, f, 0x1B)
	fmt.Println("           port",port,"              ")
	fmt.Println(" 	  -version-0.01-            ")
	fmt.Println("------------------------------------")

	// 中间件初始化
	app.Initmiddeware()

	// 注册路由
	// index/Index index=indexcontroller Index=Index
	app.SetMiddlewareFunc("index/Index",middlewarefuncs.Index_middleware)

	// 控制器注册
	app.AutoRouter(&controller.IndexController{})
	app.ServerRun_Http(":8082")
}
