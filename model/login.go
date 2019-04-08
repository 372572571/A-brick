package usermodel

import (
	model "A-brick/model/basemodel"
	"fmt"

	"github.com/372572571/Exercise/chanrpc"
)

// User 用户表
type User struct {
	ID      uint   // 用作登陆账号
	Name    string // 昵称
	Account string // 用在密码
	// CreateAt time.Time
}

// LoginModel ...
type LoginModel struct {
	*model.Model
}

// Init 初始化
func (m *LoginModel) Init() {
	m.Model = model.NewModel()
	m.registered()
	m.Model.Server.Run()

}

// 注册模型方法提供服务调用 注册函数返回值 ([]interface{}, error) 或 (error)
func (m *LoginModel) registered() {
	m.Model.Server.Registered("login", func(ages []interface{}) ([]interface{}, error) {
		// 尝试数据库添加操作
		db, _ := m.GetDb()
		if db == nil {
			return nil, model.ErrCallFail // 数据库打开失败
		}
		defer db.Close() // 关闭链接
		// db.Create(&User{ID: 103, Name: "liuyonglong", Account: "123456789"})
		l := model.TableSize(db, "users")
		fmt.Println(l, "表长度")
		return nil, nil
	})
}

// Go 调用模型方法
func (m *LoginModel) Go(key interface{}, args []interface{}) *chanrpc.Result {
	return m.Model.Server.Fast(key, args)
}
