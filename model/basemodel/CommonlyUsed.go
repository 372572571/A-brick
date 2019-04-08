package model

import (
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
)

const (
	s    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	n    = "0123456789"
	sign = "#$-+_"
)

// 常用方法

// TableSize ... 获取表长度
func TableSize(con *gorm.DB, tableName string) int {
	var res int
	con.Table(tableName).Count(&res)
	return res
}

// RandPWD 随机密码生成
// number 生成多少位
func RandPWD(number int) string {
	var res string
	rand.Seed(time.Now().Unix()) // 设置时间不然每次随机的都是一样的数.
	for i := 0; i < number; i++ {
		switch rand.Intn(3) {
		case 0:
			res += string(s[rand.Intn(len(s))])
			break
		case 1:
			res += string(n[rand.Intn(len(n))])
			break
		case 2:
			res += string(sign[rand.Intn(len(sign))])
			break
		}
	}
	return res
}
