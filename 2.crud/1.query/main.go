package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gormdemo/conn"
	"gormdemo/model"
)

var db *gorm.DB
const TIMEFORMATTER  = "2006-01-02 15:04:05"

func main() {
	db = conn.Conn()
	db.Set("gorm:table_options","ENGINE=InnoDB").AutoMigrate(&model.AdminUser{})

	defer db.Close()

	getOne()


	fmt.Println("success")
}

/**
	查询单条记录
 */
func getOne()  {
	// 声明一个对象
	var adminuser model.AdminUser
	// 1.执行First()，就是执行SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL ORDER BY `jh_admin_users`.`id` ASC LIMIT 1
	db.First(&adminuser)
	fmt.Println("First",adminuser)
	//对日期进行格式化处理
	fmt.Println(adminuser.CreatedAt.Format(TIMEFORMATTER))
	// 2.执行Last()，就是执行SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL ORDER BY `jh_admin_users`.`id` DESC LIMIT 1
	db.Last(&adminuser)
	fmt.Println("Last",adminuser)
}