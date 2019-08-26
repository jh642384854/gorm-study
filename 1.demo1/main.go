package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func main() {
	mysqlSource := "root:jianghua@tcp(127.0.0.1:3306)/godb?charset=utf8"
	db,err := gorm.Open("mysql",mysqlSource)
	if err != nil{
		log.Print(err.Error())
	}
	gorm.DefaultTableNameHandler = DefaultTableNameSet
	//自动迁移功能。注意：自动迁移仅仅会创建表，缺少列和索引，并且不会改变现有列的类型或删除未使用的列以保护数据。
	//db.AutoMigrate(&AdminUser{})
	//这个在自动创建这个表的时候，添加了一些额外的属性。这里就是将这个表的存储引擎设置为Innodb。注意，b.Set("gorm:table_options","ENGINE=InnoDB")并不是全局设置，不能单独使用
	db.Set("gorm:table_options","ENGINE=InnoDB").AutoMigrate(&AdminUser{})
	db.Set("gorm:table_options","ENGINE=InnoDB").AutoMigrate(&Link{})

	defer db.Close()
	fmt.Println("gorm success")
}
//adminuser模型
type AdminUser struct {
	gorm.Model
	Username string
	Age int
}
//link模型
type Link struct {
	gorm.Model
	//Linkid int64 `gorm:"primary_key"` // 设置Linkid为主键。如果该struct是继承gorm.Model的话，又在该struct中定义了一个字段的属性也是主键，那这个模型不会被自动创建
	Name string
}
//设置表名，可以通过给AdminUser struct类型定义TableName函数，返回一个字符串作为表名。这个优先级会高于gorm.DefaultTableNameHandler的配置
func (adminuser AdminUser) TableName() string  {
	return "jh_adminuser"
	//return "zh_adminuser"
}

//定义一个全局函数用来声明模型的前缀(即表前缀)
func DefaultTableNameSet (db *gorm.DB, defaultTableName string) string {
	return "jh_"+defaultTableName
}