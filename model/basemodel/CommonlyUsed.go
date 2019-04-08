package model

import (
	"github.com/jinzhu/gorm"
)

// 常用方法

// TableSize ... 获取表长度
func TableSize(con *gorm.DB, tableName string) int {
	var res int
	con.Table(tableName).Count(&res)
	return res
}
