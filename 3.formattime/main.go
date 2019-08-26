package main

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"fmt"
	"time"
	"log"
)

const TIMEFORMATTER  = "2006-01-02 15:04:05"
func main() {

	mysqlSource := "root:jianghua@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=true" //&parseTime=true&loc=Local
	db,err := gorm.Open("mysql",mysqlSource)
	if err != nil{
		log.Print(err.Error())
	}
	gorm.DefaultTableNameHandler = DefaultTableNameSet
	db.Set("gorm:table_options","ENGINE=InnoDB").AutoMigrate(&Product{})
	db.LogMode(true)
	defer db.Close()

	/*
	product1 := Product{Name:"apple"}
	db.Create(&product1)
	*/
	var product Product
	db.First(&product)
	jsondata,err := json.Marshal(product)
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println(string(jsondata))

	fmt.Println(product.CreatedAt.Marsha1JSON())

	fmt.Println(product.CreatedAt.Value())
}

type Product struct {
	BaseModel
	Name string
}

//使用自定义的日期格式
type MyTime struct {
	time.Time
}

//定义模型基类
type BaseModel struct {
	//gorm.Model
	ID        uint `gorm:"primary_key"`
	CreatedAt MyTime
	UpdatedAt MyTime
}
/**
	注意，这些方法并不是我想的那样，在进行对象进行JSON输出的时候，会自动调用Marsha1JSON()方法，进行这个函数处理。
	同样在获取某个字段的时候，也会自动调用Value()方法，这些都不是的，而是需要在获取该对象(MyTime)后，在调用对应的方法。
 */
func (t *MyTime) Marsha1JSON() (string,error) {
	formatted := fmt.Sprintf("\"%s\"",t.Format(TIMEFORMATTER))
	return formatted,nil
}

func (t *MyTime) Value() (driver.Value,error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano(){
		return nil,nil
	}
	return t.Time.Format(TIMEFORMATTER),nil
}

func (t *MyTime) Scan(v interface{}) error {
	value,ok := v.(time.Time)
	if ok{
		*t = MyTime{
			Time:value,
		}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func DefaultTableNameSet (db *gorm.DB, defaultTableName string) string {
	return "jh_"+defaultTableName
}