package common

import (
	"fmt"
	"server/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := Conf.DB.Host
	port := Conf.DB.Port
	username := Conf.DB.Username
	password := Conf.DB.Password
	database := Conf.DB.Database
	charset := Conf.DB.Charset
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	db.AutoMigrate(&models.User{}, &models.Role{})
	DB = db
	return db
}
