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

func (user *UserModel) BeforeCreate(org *gorm.DB) error {
	fmt.Println("before create hook function")
	// return errors.New("New Error Create By Hook Function")
	return nil
}

func main() {
	DB := Link()
	// DB = migrate(DB) // 在 gorm_new_db 数据库中，新建 user_models 表

	/* ----------------向表中插入记录操作*/

	/* 方式 1，直接插入*/
	// err := DB.Create(&UserModel{
	// 	Age:   24,
	// 	Name:  "haolan",
	// 	Email: "haolanxu@gmail.com",
	// }).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }

	/* 方式 2，回填式插入*/
	// user := UserModel{
	// 	Name:  "chenxi",
	// 	Age:   23,
	// 	Email: "xixi@gmail.com",
	// }
	// err := DB.Create(&user).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(user.Id, user.Name, user.Age, user.CreatedAt)

	/* 方式 3，批量插入*/
	// var userList = []UserModel{
	// 	{Name: "liruixi", Age: 22, Email: "ruirui@gmail.com"},
	// 	{Name: "liubing", Age: 18, Email: "bingbing@gmail.com"},
	// }
	// err := DB.Create(&userList).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("批量插入成功")

	/* 测试插入钩子函数——BeforeCreate*/
	// err := DB.Create(&UserModel{
	// 	Age:   2,
	// 	Name:  "baby",
	// 	Email: "bb@gmail.com",
	// }).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }

	/* ----------------查询表中记录操作*/
	// var userList []UserModel

	/* 方式 1，查询全部*/
	// DB.Find(&userList)
	// fmt.Println((userList))

	/* 方式 2，在全部记录中找出符合条件的*/
	// DB.Find(&userList, "id > ?", 4)
	// fmt.Println((userList))

	// var user UserModel
	/* 方式 3，查询一个*/
	// DB.Take(&user)
	// fmt.Println(user)

	/* 方式 4，按照主键查询一个*/
	// DB.Debug().Take(&user, 5) // 使用 .Debug() 查看 gorm 生成的 SQL 语句
	// fmt.Println(user)

	/* 方式 5，从前向后查询第一个符合条件的记录*/
	// DB.Debug().First(&user, 3)
	// fmt.Println(user)

	/* 方式 6，从后向前查询第一个符合条件的记录*/
	// DB.Last(&user, 4)
	// fmt.Println(user)

	/* Take 、 First 和 Last 会携带主键，如果不存在会报错*/
	// user.Id = 200
	// err := DB.Take(&user).Error
	// if err == gorm.ErrRecordNotFound {
	// 	log.Fatal("RECORD NOT FOUND")
	// }
	// fmt.Println(user)

	/* 使用 Find 来替代，不会报错*/
	// err := DB.Limit(1).Find(&user, 200).Error
	// if err == gorm.ErrRecordNotFound {
	// 	log.Fatal("RECORD NOT FOUND")
	// }
	// fmt.Println(user)

	/* ----------------更新表中记录操作*/

	/* 用 Save 创建新的记录*/
	// var user UserModel
	// user.Name = "liuyonghui"
	// user.Email = "yonghui@gmail.com"
	// DB.Save(&user)
	// fmt.Println(user)

	/* 用 Save 更新已有的记录*/
	// var user UserModel
	// user.Id = 2
	// user.Age = 24
	// user.Name = "xuhaolan"
	// user.Email = "lanlan@qq.com"
	// user.CreatedAt = time.Now()
	// DB.Save(&user)
	// fmt.Println(user)

	/* 用 Update 更新记录，需要指定主键 Id*/
	// var user = UserModel{Id: 2}
	// DB.Model(&user).Update("Age", 2)
	// fmt.Println(user)

	/* 向 Updates 中传入 struct 来更新*/
	// var user = UserModel{Id: 5}
	// DB.Model(&user).Updates(UserModel{
	// 	Name: "LB",
	// 	Age:  24,
	// })
	// fmt.Println(user)

	/* 向 Updates 中传入 mapping 来更新*/
	// var user = UserModel{Id: 4}
	// DB.Model(&user).Updates(map[string]any{
	// 	"name":  "ruirui",
	// 	"age":   1,
	// 	"email": "ruirui@qq.com",
	// })
	// fmt.Println(user)

	/* 用 UpdateColumn 更新记录，功能和 Update 一样，不会触发钩子函数*/
	// var user = UserModel{Id: 2}
	// DB.Model(&user).UpdateColumn("Age", 13)
	// fmt.Println(user)

	/* gorm.Expr*/
	// var user = UserModel{Id: 2}
	// DB.Model(&user).Update("Age", gorm.Expr("Age + ?", 10))
	// fmt.Println(user)

	/* ----------------删除表中记录操作*/

	/* 硬删除，方式 1*/
	// var user = UserModel{Id: 7}
	// DB.Delete(&user)
	// fmt.Println(user)

	/* 硬删除，方式 2*/
	// DB.Delete(&UserModel{}, 6)

	/* 按照给定的条件删除，方式 3*/
	DB.Delete(&UserModel{}, "name = ?", "ruirui")

	/* 批量删除，方式 4*/
	DB.Delete(&UserModel{}, []int{3, 5})
}
