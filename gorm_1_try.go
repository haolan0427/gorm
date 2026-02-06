package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定一个 struct 来记录表 users 中的字段
type User struct {
	Id   int
	Name string
	age  int
}

// 定一个 struct 来记录表 customers 中的字段
type Customer struct {
	Id   int
	Name string
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:7809)/gorm_new_db"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // 打开数据库，并设定配置
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
}
