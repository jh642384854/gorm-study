package main

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"gormdemo/conn"
	"gormdemo/model"
)
var db *gorm.DB


func main() {
	db = conn.Conn()
	db.Set("gorm:table_options","ENGINE=InnoDB").AutoMigrate(&model.AdminUser{})

	defer db.Close()
	// 1.删除单条记录
	// singleDelete()
	// 2.批量删除记录
	 batchDelete()
	// 3.物理删除记录(这个不管当前被删除的记录是否被软删除，都会被直接删掉)
	//physicalDelete()

}
/**
	单一删除
 */
func singleDelete()  {
	// 1.根据查询的记录来删除
	var adminuser model.AdminUser
	//db.Where("id=?",1).Find(&adminuser)
	// 执行的SQL语句：UPDATE `jh_admin_users` SET `deleted_at`='2019-08-27 15:58:49'  WHERE `jh_admin_users`.`deleted_at` IS NULL AND `jh_admin_users`.`id` = 1
	//db.Delete(&adminuser)

	// 2.根据查询条件来进行删除
	// 上面的删除一条记录执行了两次查询操作，使用这种方式，只需要执行一次就好了。
	// 执行的SQL语句:UPDATE `jh_admin_users` SET `deleted_at`='2019-08-27 16:01:29'  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((`jh_admin_users`.`username` = 'sipotian'))
	//db.Delete(&adminuser,model.AdminUser{Username:"sipotian"})

	// 3.根据主键ID来进行删除
	// 如果是依据主键来进行删除，Delete()的第二个参数还可以直接写主键ID
	db.Delete(&adminuser,3)
}

/**
	批量删除
 */
func batchDelete()  {
	// 删除status值为0的记录
	// 执行SQL：UPDATE `jh_admin_users` SET `deleted_at`='2019-08-27 16:07:00'  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((status = 0))
	db.Where("status = ?",0).Delete(model.AdminUser{})
	fmt.Println(db.RowsAffected)
}

/**
	物理删除
	需要使用Unscoped()方法
 */
func physicalDelete()  {
	// 永久的删除记录
	// 执行的SQL语句：DELETE FROM `jh_admin_users`  WHERE (`jh_admin_users`.`id` = 3)
	db.Unscoped().Delete(model.AdminUser{},3)
} 