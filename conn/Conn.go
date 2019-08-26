package conn

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)
func Conn() *gorm.DB {
	/**
		这个数据库链接字符串，后面的一些请求参数也很重要：
		charset：设置数据库的字符集
		parseTime：设置为true，这个就会在被定义的字段是time.Time类型在获取数据数据的时候会解析存储在数据库里面的记录。
		loc：Local，这个用来定义mysql的时区，如果不声明的话，写入到数据表的记录时间就会又缺失。
	 */
	mysqlSource := "root:jianghua@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=true&loc=Local"
	db,err := gorm.Open("mysql",mysqlSource)
	if err != nil{
		log.Print(err.Error())
	}
	//设置表前缀
	gorm.DefaultTableNameHandler = DefaultTableNameSet
	//是否开启sql语句调试
	db.LogMode(true)
	return db
}

func DefaultTableNameSet (db *gorm.DB, defaultTableName string) string {
	return "jh_"+defaultTableName
}