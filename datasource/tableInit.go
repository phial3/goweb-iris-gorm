package datasource

import (
	"goweb-iris-gorm/models"
)

// 初始化表 如果不存在该表 则自动创建

func CreateTable() {
	GetDB().AutoMigrate(
		&models.User{},
		&models.Book{},
	)
}
