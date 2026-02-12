package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserModel struct {
	Id        int
	Name      string `gorm:"not null;unique"`
	Age       int    `gorm:"default:18"`
	Email     string `gorm:"size:32"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func Link() *gorm.DB { // 封装数据库连接操作
	dsn := "root:root@tcp(127.0.0.1:7809)/gorm_new_db?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func migrate(DB *gorm.DB) *gorm.DB { // 封装数据库表迁移操作
	err := DB.AutoMigrate(&UserModel{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("表迁移成功")
	return DB
}

func (user *UserModel) BeforeDelete(db *gorm.DB) (err error) {
	fmt.Println("hook function")
	return nil
}

func main() {
	DB := Link()
	// DB = migrate(DB)

	/* 软删除，会在 DeletedAt 加上删除时间，但并没有删除*/
	// DB.Delete(&UserModel{}, "age <= ?", 10)

	/* 使用 Take ，查询不到被软删除了的记录*/
	// var user UserModel
	// DB.Debug().Take(&user, "age <= ?", 10)
	// fmt.Println(user)

	/* 使用 Unscoped.Take 查询得到被软删除了的记录*/
	// var user UserModel
	// DB.Unscoped().Take(&user, "age <= ?", 10)
	// fmt.Println(user)

	/* 彻底删除 被软删除了的记录*/
	DB.Unscoped().Delete(&UserModel{}, "age = ?", 10)

}
