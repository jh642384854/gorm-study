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

//如果你想在BeforeCreate hook 中修改字段的值，可以使用scope.SetColumn
func (adminuser AdminUser) BeforeCreate(scope *gorm.Scope) error {
	uuidV := uuid.NewV4().String()
	if err := scope.SetColumn("uuid",uuidV);err != nil{
		log.Println(err.Error())
	}
	return nil
}