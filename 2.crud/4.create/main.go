package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/twinj/uuid"
	"log"
	"time"
)
func main() {
	mysqlSource := "root:jianghua@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=true&loc=Local" //parseTime=true&loc=America%2FChicago
	db,err := gorm.Open("mysql",mysqlSource)
	if err != nil{
		log.Print(err.Error())
	}
	gorm.DefaultTableNameHandler = DefaultTableNameSet
	db.Set("gorm:table_options","ENGINE=InnoDB").AutoMigrate(&AdminUser{})

	defer db.Close()
	//插入数据
	//insert(db)
	//获取数据
	var adminuser AdminUser
	db.First(&adminuser)
	fmt.Println(adminuser)

	var article Article
	db.Find(&article)
	fmt.Println(article)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}

func insert(db *gorm.DB)  {
	user1 := AdminUser{Username:"sunli",Status:1,LoginFails:10}
	//user2 := AdminUser{Username:"lisi",Age:20}
	//user3 := AdminUser{Username:"wangmaz",Age:20}
	// 使用Create()函数，返回DB对象。函数需要传递引用地址。
	db.Create(&user1)
	//fmt.Println(db.RowsAffected,user1.ID)
	// 需要先创建模型的对象，在插入数据后通过NewRecord返回值判断是否插入成功，如果成功则返回false
	boolV := db.NewRecord(&user1)
	fmt.Println(boolV)
	if !boolV {
		fmt.Println("success")
	}else{
		fmt.Println("fail")
	}
}

type Article struct {
	ID int
	Title string
	Author string
	Status int
	Views int
}

type AdminUser struct {
	gorm.Model
	Uuid string
	Username string
	Age int `gorm:"default:'25'"`   // 这里设置，会将表结构的该字段值设置为25，后面在进行插入数据时候，不指定该字段值，就会用默认值
	Status int
	LoginFails int
}

//如果你想在BeforeCreate hook 中修改字段的值，可以使用scope.SetColumn
func (adminuser AdminUser) BeforeCreate(scope *gorm.Scope) error {
	uuidV := uuid.NewV4().String()
	fmt.Println(uuidV)
	if err := scope.SetColumn("uuid",uuidV);err != nil{
		log.Println(err.Error())
	}
	return nil
}

func DefaultTableNameSet (db *gorm.DB, defaultTableName string) string {
	return "jh_"+defaultTableName
}