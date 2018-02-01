package app

import (
	"net/http"
)

// 禁止URL执行的方法
var shieldfuncs = make(map[string]int)



// 添加中间件规则方法
// [控制器/方法]func()中间件方法
var middlewareFunc=make(map[string]func(w http.ResponseWriter,r *http.Request))

// 初始化 middleware
func Initmiddeware(){
	SetShieldfuncs() // 屏蔽App方法，防止URL调用
}

// 注册中间件
// cm 控制器/方法
// func 绑定的中间件方法
func SetMiddlewareFunc(cm string,f func(w http.ResponseWriter,r *http.Request)){
	middlewareFunc[cm]=f
	return
}

// 设置 禁止URL执行的方法
func SetShieldfuncs(){
	shieldfuncs=appLimit()
}

// 判断某个控制器下的方法是否已经注册过中间件
func IsSetMiddleware(cm string)bool{
	if _,ok:=middlewareFunc[cm];ok{
		return false
	}
	return true
}

// 判断方法是否被屏蔽 屏蔽返回false
func IsShieldfunc(s string)bool{
	if _,ok:=shieldfuncs[s];ok {
		return false
	}
	return true
}




