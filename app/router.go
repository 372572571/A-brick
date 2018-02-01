package app

import (
	"reflect"
	"strings"
	"net/http"
)

var webRoute map[string]reflect.Type = make(map[string]reflect.Type)
// 添加路由
func router(pattern string, t reflect.Type) {
	webRoute[strings.ToLower(pattern)] = t
}
// 储存反射信息到路由
func Router(pattern string, app IApp) {
	refV := reflect.ValueOf(app)
	refT := reflect.Indirect(refV).Type()
	router(pattern, refT)
}
// 注册路由
func AutoRouter(app IApp) {
	refV := reflect.ValueOf(app)
	refT := reflect.Indirect(refV).Type()
	refName := strings.TrimSuffix(strings.ToLower(refT.Name()), "controller")// 去掉controller
	router(refName, refT)//储存
}

var Static map[string]string = make(map[string]string)// 静态服务

func serveStatic(w http.ResponseWriter, r *http.Request) bool {
	for prefix, static := range Static {// 循环注册静态服务
		if strings.HasPrefix(r.URL.Path, prefix) {
			file := static + r.URL.Path[len(prefix):]
			http.ServeFile(w, r, file)
			return true
		}
	}

	return false
}