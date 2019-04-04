package model

import (
	"github.com/372572571/Exercise/chanrpc"
)

// ModelPool 存放服务的数据库rpc服务
// 模型池 根据控制器的名称存放
var ModelPool map[string]IModel = make(map[string]IModel)

// IModel 服务结构
type IModel interface {
	Init()                                                  // 初始化
	Go(key interface{}, args []interface{}) *chanrpc.Result // 调用模型方法
}

// Model ...
type Model struct {
	server *chanrpc.Server
}

// NewModel ... 创建模型
func NewModel() *Model {
	var s = chanrpc.NewServer(50)
	return &Model{server: s}
}

// ModelPoolAdd ...
func ModelPoolAdd(key string, model IModel) {
	ModelPool[key] = model
}

// ModelPoolRun 初始化池模型
func ModelPoolRun() {
	for _, value := range ModelPool {
		value.Init()
	}
}
