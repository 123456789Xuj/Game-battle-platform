package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL() error {
	dns := "root:xjl666nbsg@tcp(127.0.0.1:3306)/landing?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_ = DB.AutoMigrate(&NewUser{})
	_ = DB.AutoMigrate(&GameRoom{})
	_ = DB.AutoMigrate(&GameRoom{})
	_ = DB.AutoMigrate(&Chess{})
	_ = DB.AutoMigrate(&League{})
	return nil

}
