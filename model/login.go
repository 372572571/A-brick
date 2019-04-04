package model

import (
	"fmt"

	"github.com/372572571/Exercise/chanrpc"
)

// LoginModel ...
type LoginModel struct {
	*Model
}

// Init 初始化
func (m *LoginModel) Init() {
	m.Model = NewModel()
	m.registered()
	m.Model.server.Run()
}

// 注册模型方法
func (m *LoginModel) registered() {
	m.Model.server.Registered("login", func(ages []interface{}) ([]interface{}, error) {
		fmt.Println(ages)
		fmt.Println("模型ok")
		return nil, nil
	})
	// fmt.Println(m.Model.server)
}

// Go 调用模型方法
func (m *LoginModel) Go(key interface{}, args []interface{}) *chanrpc.Result {
	fmt.Println("模型方法调用")
	return m.Model.server.Fast(key, args)
}
