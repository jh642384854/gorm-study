package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gormdemo/conn"
	"gormdemo/model"
)

/**
	原生SQL操作。就是直接执行SQL语句
 */

var db *gorm.DB

func main() {
	db = conn.Conn()
	db.Set("gorm:table_options","ENGINE=InnoDB").AutoMigrate(&model.AdminUser{})

	defer db.Close()

/*
	// 1.更新操作
	sqlStr := "UPDATE jh_admin_users SET age = ?,login_fails = ? WHERE id = ?"
	if err := db.Exec(sqlStr,32,1,1).Error; err != nil{
		fmt.Println("Execute Sql error:",err.Error())
	}
*/
	// 2.查询操作，Raw()和Scan()函数的配合使用
	var adminuser model.AdminUser
	sqlStr2 := "SELECT * FROM jh_admin_users WHERE id = ?"
	if err := db.Raw(sqlStr2,1).Scan(&adminuser).Error; err != nil{
		fmt.Println("Sql Error",err.Error())
	}else{
		fmt.Println(adminuser)
	}

	// 3.Row()和Rows()方法的使用
	var username,uuid string
	//这里需要注意一下，Select()函数的第一个参数，查询的字段都需要写在一个字符串中，不能这样写Select("username","uuid")这样是不对的。
	row := db.Table("jh_admin_users").Where("id = ?", 1).Select("username, uuid").Row() // (*sql.Row)
	if err := row.Scan(&username, &uuid);err != nil{
		fmt.Println("Sql Error",err.Error())
	}
	fmt.Println(username,uuid)
	fmt.Println()

	// 4.多条记录指定字段操作，形式一.推荐用这种方式，为什么呢？因为这种会自动带上deleted_at属性的判断。如果用下面的一种方式，很有可能就会遗忘这个字段的判断，这样获取的数据可能就是不正确的。
	// 执行的SQL:SELECT username,uuid FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((id > 3))
	rows,err := db.Model(&model.AdminUser{}).Where("id > ?",3).Select("username,uuid").Rows()
	if err != nil{
		fmt.Println("Sql Error",err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&username,&uuid);err != nil{
			continue
		}
		fmt.Println(username,uuid)
	}
	fmt.Println()

	// 4.多条记录指定字段操作，形式二。使用Raw()函数
	sqlStr3 := "SELECT username,uuid FROM `jh_admin_users`  WHERE id > ?"
	// 执行的SQL:SELECT username,uuid FROM `jh_admin_users`  WHERE id > 5
	rows2,err := db.Raw(sqlStr3,5).Rows()
	defer rows2.Close()
	for rows2.Next() {
		if err := rows2.Scan(&username,&uuid);err != nil{
			continue
		}
		fmt.Println(username,uuid)
	}
	fmt.Println()

	// 5.多条记录对象操作，这里使用的是gorm.ScanRows()方法
	// 执行的SQL：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((id > 3))
	row3,err := db.Model(&model.AdminUser{}).Where("id > ?",3).Rows()
	if err != nil{
		fmt.Println("Sql Error",err.Error())
	}
	defer row3.Close()
	for row3.Next() {
		var adminuser model.AdminUser
		if err := db.ScanRows(row3,&adminuser); err != nil{
			continue
		}
		fmt.Println(adminuser)
	}
}
