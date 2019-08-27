package model

import (
	"github.com/jinzhu/gorm"
	"github.com/twinj/uuid"
	"log"
)

type AdminUser struct {
	gorm.Model
	Uuid string
	Username string
	Age int `gorm:"default:'25'"`   // 这里设置，会将表结构的该字段值设置为25，后面在进行插入数据时候，不指定该字段值，就会用默认值
	Status int
	LoginFails int
}
// 为当前模型定义Callbacks，简单的说，就是为当前模型定义了一些hook(钩子)
// 需要注意的是，这些钩子函数的形参是没有约定的，可以不传递，可以传递*grom.Scope、*gorm.DB等等对象，这个根据需要来定。

// 如果你想在BeforeCreate hook 中修改字段的值，可以使用scope.SetColumn
func (adminuser AdminUser) BeforeCreate(scope *gorm.Scope) error {
	uuidV := uuid.NewV4().String()
	if err := scope.SetColumn("uuid",uuidV);err != nil{
		log.Println(err.Error())
	}
	return nil
}


func (adminuser *AdminUser) BeforeSave()  {
	
}

func (adminuser *AdminUser) AfterCreate()  {

}

func (adminuser *AdminUser) AfterSave()  {

}

func (adminuser *AdminUser) BeforeDelete()  {

}

func (adminuser *AdminUser) AfterDelete()  {

}

func (adminuser *AdminUser) AfterFind()  {

}