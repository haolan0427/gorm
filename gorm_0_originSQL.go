package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "gorm.io/driver/mysql"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:7809)/gorm_new_db"
	db, err := sql.Open("mysql", dsn) // 用这个 err 来判断 dsn 是否错误
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil { // 用这个 err 来判断是否连接上数据库
		log.Fatal(err)
	}
	fmt.Println(db)

	/* Exec 函数，建表*/
	// _, err = db.Exec("create table customers (id bigint unsigned auto_increment primary key, name varchar(50) not null default '' comment '用户昵称') engine = innodb default charset = utf8mb4 collate = utf8mb4_unicode_ci")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	/* Exec 函数，插入数据*/
	// _, err = db.Exec("insert into customers (name) values ('zhangsan'), ('lisi'), ('wangwu')")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	/* Query ，查询多行*/
	// rows, err := db.Query("select name, id from customers")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for rows.Next() {
	// 	var id int
	// 	var name string
	// 	err = rows.Scan(&name, &id) //这里的参数顺序要和 db.Query 中的查询 表属性 顺序一致
	// 	fmt.Println(id, name, err)
	// }

	/* QueryRow ，查询单行*/
	var id int
	var name string
	err = db.QueryRow("select id, name from customers").Scan(&id, &name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id, name)

}
