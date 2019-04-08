package model

import (
	"fmt"

	"github.com/372572571/Exercise/chanrpc"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	Server *chanrpc.Server
}

// NewModel ... 创建模型
func NewModel() *Model {
	var s = chanrpc.NewServer(50)
	return &Model{Server: s}
}

// GetDb 获取数据库链接
func (m *Model) GetDb() (*gorm.DB, error) {
	var db, err = gorm.Open("mysql", "root_lyl:LIUYONGLONGLOVEZS85988984a@(101.132.97.64:3306)/GOGAME?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("创建链接失败")
		return nil, err
	}
	return db, nil
}

// PoolAdd ...
func PoolAdd(key string, model IModel) {
	ModelPool[key] = model
}

// PoolRun 初始化池模型
func PoolRun() {
	for _, value := range ModelPool {
		value.Init()
	}
}
