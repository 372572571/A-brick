package app

import (
	"sync"
	"net/http"
	"strings"
	"reflect"
)

const (
	defController="index"
	defMethod="Index"
)

type handler struct {
	p sync.Pool
}
// 创建一个myHandler
func newHandler() *handler{
	myh:=&handler{}
	myh.p.New= func() interface{} {
		return &Context{}// 保存，请求和请求回应
	}
	return myh
}

// 实现serveHTTP
// type Handler interface {
// ServeHTTP(ResponseWriter, *Request)

func(h *handler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	// 静态处理
	if serveStatic(w, r) {
		return
	}
	// fmt.Println(r.Method)
	// fmt.Println(r.Cookies())

	ctx := h.p.Get().(*Context) // get获得interface数据.(*context)强制转换
	defer h.p.Put(ctx)
	ctx.Config(w, r) // 储存数据

	controllerName, methodName := h.findControllerInfo(r) // 获得 控制器名称，方法名称
	// fmt.Println(controllerName,methodName)
	controllerT, ok := webRoute[controllerName] // 查找是否注册控制器，存在返回一个refkkect.type
	if !ok {
	http.NotFound(w, r)
		return
	}
	refV := reflect.New(controllerT) // 根据路由注册的信息，创建一个新的结构
	method := refV.MethodByName(methodName) // 查找是否有这个方法
	if !method.IsValid() {
	http.NotFound(w, r)
		return
	}
	// 中间件.判断.存在则执行 注册的中间件方法
	if ismiddleware,ok:=middlewareFunc[controllerName+"/"+methodName];ok{
		ismiddleware(ctx.w,ctx.r)
	}
	controller := refV.Interface().(IApp) // 返回接口,强制转换成IApp
	controller.Init(ctx) // 通过接口调用方法.初始化controller
	method.Call(nil) //执行对应的方法

}
// 查询对应C,M 信息
// 返回控制器，方法名称
func(h *handler)findControllerInfo(r *http.Request)(controllerName string ,methodName string){
	path := r.URL.Path
	if strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}
	pathInfo := strings.Split(path, "/")

	controllerName = defController
	if len(pathInfo) > 1 {
		controllerName = pathInfo[1]
	}

	methodName = defMethod
	if len(pathInfo) > 2 {
		// strings.Title(strings.ToLower(pathInfo[2]))
		// 方法名称规范,首字母大写,其余小写
		methodName = strings.Title(strings.ToLower(pathInfo[2]))
		if !IsShieldfunc(methodName){
			methodName=defMethod
			controllerName = defController
		}
	}

	// fmt.Println(controllerName,methodName)
	return
}
