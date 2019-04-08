package usermodel

import (
	model "A-brick/model/basemodel"

	"github.com/372572571/Exercise/chanrpc"
)

// User 用户表
type User struct {
	ID      uint   // 用作登陆账号
	Name    string // 昵称
	Account string // 用在密码
	Status  int
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
	// 创建随机用户服务
	m.Model.Server.Registered("create_rand_user", m.createRandUser)
}

// Go 调用模型方法
func (m *LoginModel) Go(key interface{}, args []interface{}) *chanrpc.Result {
	return m.Model.Server.Fast(key, args)
}

// createRandUser 根据数据库表长度随机创建一个用户
// 返回结果 [0]interface{}
func (m *LoginModel) createRandUser(args []interface{}) ([]interface{}, error) {
	// 获取数据库长度
	db, _ := m.GetDb()
	if db == nil {
		return nil, model.ErrorLinkFail
	}
	defer db.Close() // 关闭链接
	var l = uint(model.TableSize(db, "users") + 50000)
	l = uint(l)
	u := &User{ID: l, Name: "Tourist", Account: model.RandPWD(15), Status: 1}
	db.Create(u)
	res := []interface{}{u}
	return res, nil
}
