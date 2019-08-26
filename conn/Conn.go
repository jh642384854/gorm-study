package conn

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)
func Conn() *gorm.DB {
	mysqlSource := "root:jianghua@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=true&loc=Local"
	db,err := gorm.Open("mysql",mysqlSource)
	if err != nil{
		log.Print(err.Error())
	}
	//设置表前缀
	gorm.DefaultTableNameHandler = DefaultTableNameSet
	//是否开启sql语句调试
	//db.LogMode(true)
	return db
}

func DefaultTableNameSet (db *gorm.DB, defaultTableName string) string {
	return "jh_"+defaultTableName
}