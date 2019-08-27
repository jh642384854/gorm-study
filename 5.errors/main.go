package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gormdemo/conn"
	"gormdemo/model"
)

/**
	错误处理
 */

var db *gorm.DB


func main() {
	db = conn.Conn()
	db.Set("gorm:table_options","ENGINE=InnoDB").AutoMigrate(&model.AdminUser{})

	defer db.Close()

	// 1.普通的错误处理
	var adminuser model.AdminUser
	if err := db.Where("id = ?",1).First(&adminuser).Error; err != nil{
		fmt.Println(err.Error())
	}else{
		fmt.Println(adminuser)
	}

	// 2.当进行链式操作的时候，可能不知道哪个地方出问题，这样就可以在链式最后调用GetErrors()，然后根据这个返回值来判断是否有错误
	var adminuser2 model.AdminUser
	errors := db.First(&adminuser2).Limit(10).Find(&adminuser2).GetErrors()
	fmt.Println(len(errors))
	for _, err := range errors {
		fmt.Println(err)
	}

	// 3.没有查到记录错误处理。通过IsRecordNotFoundError()函数来判断错误
	var adminuser3 model.AdminUser
	// 方式一：
	if err := db.Where("id = ?",10).First(&adminuser3).Error; gorm.IsRecordNotFoundError(err){
		fmt.Println("没有找到相应记录")
	}else{
		fmt.Println(adminuser3)
	}
	// 方式二：使用RecordNotFound()函数
	var adminuser4 model.AdminUser
	if db.Where("id = ?",10).First(&adminuser4).RecordNotFound(){
		fmt.Println("没有找到相应记录")
	}else{
		fmt.Println(adminuser4)
	}
}
