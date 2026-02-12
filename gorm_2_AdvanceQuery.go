package main

import (
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

func main() {
	DB := Link()

	/* 使用 Where 进行普通查询，要在 Take 等之前*/
	// var user UserModel
	// DB.Where("age < 12").Take(&user)
	// fmt.Println(user)

	/* 使用 Where 进行 struct 查询*/
	// var user UserModel
	// DB.Where(UserModel{
	// 	Name: "xuhaola",
	// 	Age:  23,
	// }).Take(&user)
	// fmt.Println(user)

	/* 使用 Where 进行 mapping 查询*/
	// var user UserModel
	// DB.Debug().Where(map[string]any{
	// 	"name": "lisi",
	// }).Take(&user)
	// fmt.Println(user)

	/* 链式调用*/
	// query := DB.Where("age > ? and id < ?", 10, 5)
	// var user UserModel
	// err := query.Where("name != ?", "xuhaolan").Take(&user).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(user)

	/* Or 查询*/
	// var userList []UserModel
	// DB.Or("id < ?", 2).Or("name = ?", "lisi").Find(&userList)
	// fmt.Println(userList)

	/* 按照主键升序查找*/
	// var userList []UserModel
	// DB.Find(&userList)
	// fmt.Println(userList)

	/* 按照 id 字段降序查找*/
	// var userList []UserModel
	// DB.Order("id desc").Find(&userList)
	// fmt.Println(userList)

	/* 按照 age 字段升序查找*/
	// var userList []UserModel
	// DB.Order("age asc").Find(&userList)
	// fmt.Println(userList)

	/* 按照 age 字段升序，如果 age 相同，再用 id 降序查找*/
	// var userList []UserModel
	// DB.Order("age asc").Order("id desc").Find(&userList)
	// fmt.Println(userList)

	/* 高级查询， Select 与 Scan*/
	// var userList []string // string 对应 UserModel.Name 的类型
	// DB.Model(UserModel{}).Select("name").Scan(&userList)
	// fmt.Println(userList)

	/* 高级查询， Pluck，集成了 Select 与 Scan 实现的高级查询*/
	// var userList []string
	// DB.Model(UserModel{}).Pluck("name", &userList)
	// fmt.Println(userList)

	/* 去重 Distinct*/
	// var ageList []int
	// DB.Model(UserModel{}).Distinct().Pluck("age", &ageList)
	// fmt.Println(ageList)

	/* 扫描到结构体中*/
	// type User struct {
	// 	A int    `gorm:"column:age"`
	// 	B string `gorm:"column:name":`
	// }
	// var userList []User
	// DB.Debug().Model(UserModel{}).Scan(&userList)
	// fmt.Println(userList)

	/* Group 与 Select count()*/
	// type User struct {
	// 	Age   int
	// 	Count int
	// }
	// var userList []User
	// DB.Debug().Model(UserModel{}).Group("age").Select("age, count(id) as count").Scan(&userList)
	// fmt.Println(userList)

	/* Group 与 Select sum()*/
	// type User struct {
	// 	Age   int
	// 	Total int
	// }
	// var userList []User
	// DB.Debug().Model(UserModel{}).Group("age").Select("age, sum(age) as total").Scan(&userList)
	// fmt.Println(userList)

	/* 使用 Limit 和 Offset 实现分页操作*/
	// var userList []UserModel
	// limit := 2
	// page := 2 // 可以设置为 1、2、3，表示第 1、2、3页
	// DB.Limit(limit).Offset((page - 1) * limit).Find(&userList)
	// fmt.Println(userList)

	/* Scopes 配合自定义的 Age10 函数*/
	// var userList []UserModel
	// DB.Debug().Scopes(Age10).Find(&userList)
	// fmt.Println(userList)

	// /* Scopes 配合自定义的 NameIn 函数*/
	// var userList []UserModel
	// DB.Debug().Scopes(NameIn("xuhaolan", "nobody", "lisi")).Find(&userList)
	// fmt.Println(userList)

	/* Raw*/
	// type User struct {
	// 	Name string
	// 	Age  int
	// }
	// var userList []User
	// DB.Debug().Raw("select name, age from user_models where age < ?", 18).Scan(&userList)
	// fmt.Println(userList)

	/* Exec*/
	DB.Debug().Exec("update user_models set age = ? where name = ?", 21, "luliu")
}

func Age10(tx *gorm.DB) *gorm.DB {
	return tx.Where("age > ?", 10)
}

func NameIn(nameList ...string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where("name in ?", nameList)
	}
}
