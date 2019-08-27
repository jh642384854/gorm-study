package main

import (
	"github.com/jinzhu/gorm"
	"gormdemo/conn"
	"gormdemo/model"
)
var db *gorm.DB


func main() {
	db = conn.Conn()
	db.Set("gorm:table_options","ENGINE=InnoDB").AutoMigrate(&model.AdminUser{})

	defer db.Close()

	// 1.更新所有字段值
	// updateAllFiled()
	// 2.只是更新设定的字段
	// updateModifyFiled()
	// 3.更新特定的字段或屏蔽某些字段
	// updateSpecifiedField()
	// 4.批量更新操作
	// batchUpdate()
	// 5.gorm.Expr()使用
	exprUpdate()
}
/**
	Save()方法。更新所有字段。这个操作在实际应用中应该会很少
	下面的操作会执行下面两条语句：
	SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((`jh_admin_users`.`id` = 1)) ORDER BY `jh_admin_users`.`id` ASC LIMIT 1
	UPDATE `jh_admin_users` SET `created_at` = '2019-08-26 07:41:46', `updated_at` = '2019-08-26 07:41:46', `deleted_at` = NULL, `uuid` = '5964f734-a5b7-42c8-998c-09166bc84a64', `username` = 'chendalei', `age` = 35, `status` = 0, `login_fails` = 0  WHERE `jh_admin_users`.`deleted_at` IS NULL AND `jh_admin_users`.`id` = 1
 */
func updateAllFiled()  {
	var adminuser model.AdminUser
	db.First(&adminuser,1)
	adminuser.Username = "chendalei"
	adminuser.Age = 35
	db.Save(adminuser)
}

/**
	Update()和Updates()方法
	Update()：只能更新单个字段值
	Updates()：这个可以更新多个字段值
	需要注意的是，还有两个方法UpdateColumn()和UpdateColumns()方法，这两个方法和Update()和Updates()方法类似，都是用来更新字段的。
	但是不同的是，Update()和Updates()方法会触发模型的BeforeUpdate, AfterUpdate方法，所以，在进行更新的时候，会自动把模型的updated_at字段值进行更新。
	如果不想这样的话，就使用UpdateColumn()和UpdateColumns()方法这两个方法
 */
func updateModifyFiled()  {
	var adminuser model.AdminUser
	db.First(&adminuser,1)
	// 1.下面是修改单个字段值
	// 执行的SQL语句：UPDATE `jh_admin_users` SET `updated_at` = '2019-08-27 11:33:09', `username` = 'chengerlei'  WHERE `jh_admin_users`.`deleted_at` IS NULL AND `jh_admin_users`.`id` = 1
	//db.Model(&adminuser).Update("username","chengerlei")

	// 2.更新多个字段，这就需要用到结构体或是map的数据结构了
	var adminuser2 model.AdminUser
	db.First(&adminuser2,2)
	// 执行的SQL语句：UPDATE `jh_admin_users` SET `login_fails` = 10, `status` = 2, `updated_at` = '2019-08-27 11:38:28', `username` = 'sipotian'  WHERE `jh_admin_users`.`deleted_at` IS NULL AND `jh_admin_users`.`id` = 2
	db.Model(&adminuser2).Updates(map[string]interface{}{"username":"sipotian","status":2,"login_fails":10})

	//结构体的使用就不做例子,上面的更新操作和下面的这个是一样的
	//db.Model(&adminuser2).Updates(model.AdminUser{Username:"sipotian",Status:2,LoginFails:10})
}
/**
	配合使用Select()和Omit()函数来实现更新或不更新某些字段.
	当使用了Omit()函数后，即便在Updates()里面指定需要更新的字段，但是也会排除Omit()指定的字段值。
 */
func updateSpecifiedField()  {
	var adminuser3 model.AdminUser
	db.First(&adminuser3,3)
	// 执行的SQL语句：UPDATE `jh_admin_users` SET `status` = 2, `updated_at` = '2019-08-27 11:47:15'  WHERE `jh_admin_users`.`deleted_at` IS NULL AND `jh_admin_users`.`id` = 3
	db.Model(&adminuser3).Omit("username","login_fails").Updates(map[string]interface{}{"username":"sipotian","status":2,"login_fails":10})
}

/**
	批量更新操作.
	在批量更新时，Callbacks机制不会运行
 */
func batchUpdate()  {
	// 执行的SQL语句：UPDATE `jh_admin_users` SET `login_fails` = 10, `status` = 1, `updated_at` = '2019-08-27 11:54:32'  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((id > 5))
	db.Model(&model.AdminUser{}).Where("id > ?",5).Updates(map[string]interface{}{"status":1,"login_fails":10})
}

/**
	gorm.Expr()使用
 */
func exprUpdate()  {
	// 执行的SQL语句： UPDATE `jh_admin_users` SET `login_fails` = login_fails + 1, `updated_at` = '2019-08-27 14:00:44'  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((id=1))
	db.Model(&model.AdminUser{}).Where("id=?",1).Update("login_fails",gorm.Expr("login_fails + ?",1))
}

