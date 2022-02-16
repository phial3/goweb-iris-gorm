package datasource

import (
	"goweb-iris-gorm/model"
)

// 初始化表 如果不存在该表 则自动创建

func CreateTable() {
	GetDB().AutoMigrate(
		&model.User{},
		&model.Book{},
	)
}
