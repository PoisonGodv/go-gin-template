package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"test_wxlogin/models"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:010413@tcp(127.0.0.1:3306)/test_go_sql?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败", err)
	} else {
		fmt.Printf("数据库连接成功: %v", db)
	}
	err2 := db.AutoMigrate(&models.UserBasic{})
	if err2 != nil {
		return
	} else {
		fmt.Printf("创建表成功: %v", db)
	}
}
