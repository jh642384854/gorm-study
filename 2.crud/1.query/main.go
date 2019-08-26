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
	//adminuser := model.AdminUser{}
	var adminuser model.AdminUser
	db.Last(&adminuser)
	fmt.Println(adminuser)
	fmt.Println(adminuser.CreatedAt.Format(TIMEFORMATTER))
}