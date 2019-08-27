package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gormdemo/conn"
	"gormdemo/model"
	"time"
)

var db *gorm.DB
const TIMEFORMATTER  = "2006-01-02 15:04:05"

func main() {
	db = conn.Conn()
	db.Set("gorm:table_options","ENGINE=InnoDB").AutoMigrate(&model.AdminUser{})

	defer db.Close()
	// 1.查询单条记录
	// getOne()
	// 2.查询多条记录
	// getAll()
	// 3.简单的where条件查询
	// simpleWhereQuery()
	// 4.高级where条件查询
	// advancedWhereQuery()
	// 5.Not查询条件
	// notWhereQuery()
	// 6.单纯使用Find()函数进行的条件查询
	// findWhereQuery()
	// 7.Or查询条件
	// orWhereQuery()
	// 8.FirstOrInit()和FirstOrCreate()函数使用，以及这两个函数配合Assign()和Attrs()函数使用
	// firstOrInitQuery()
	// 9.Select()函数使用
	// selectQuery()
	// 10.Order()、Limit()、Offset()函数使用
	// pageQuery()
	// 11.获取记录总数
	// countQuery()
	// 12.使用指定的数据表来进行查询
	// tableQuery()
	// 13.Scan()函数使用
	// scanQuery()
	// 14.Pluck()函数使用
	// pluckQuery()
	// 15.Scopes()函数使用
	scopesQuery()

	fmt.Println("success")
}

/**
	查询单条记录

	注意，不管是带条件或是没有带条件的查询，都需要配合Find()、First()、Last()这些函数来实现，这三个函数的应用区别：
	①、查询单条记录多用First()或是Last()函数
	②、查询多条记录，就需要配合Find()函数
 */
func getOne()  {
	// 声明一个对象
	var adminuser model.AdminUser
	// 1.执行First()，就是执行SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL ORDER BY `jh_admin_users`.`id` ASC LIMIT 1
	db.First(&adminuser)
	fmt.Println("First",adminuser)
	//对日期进行格式化处理
	fmt.Println(adminuser.CreatedAt.Format(TIMEFORMATTER))
	// 2.执行Last()，就是执行SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL ORDER BY `jh_admin_users`.`id` DESC LIMIT 1
	var adminuser2 model.AdminUser
	db.Last(&adminuser2)
	fmt.Println("Last",adminuser2)
	// 3.查询指定ID的记录，就是执行SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((`jh_admin_users`.`id` = 2)) ORDER BY `jh_admin_users`.`id` ASC LIMIT 1
	var adminuser3 model.AdminUser
	db.First(&adminuser3,2)
	fmt.Println("First ID 2",adminuser3)
	// 4.查询指定ID的记录，就是执行SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((`jh_admin_users`.`id` = 2)) ORDER BY `jh_admin_users`.`id` DESC LIMIT 1
	var adminuser4 model.AdminUser
	db.Last(&adminuser4,2)
	fmt.Println("First ID 2",adminuser4)
	// 当使用First()和Last()函数进行查询的时候，传递第二个参数，其实这两个函数都可以随意使用。因为所达到的效果是一样的。
}

/**
	查询所有记录
 */
func getAll()  {
	var adminusers []model.AdminUser
	// 直接调用Find()方法，就是执行SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL  
	db.Find(&adminusers)
	fmt.Println(adminusers)
}

/**
	最简单的带条件查询
	下面的示例都是只提供了一个参数，如果又多个参数，就写多个参数就好了
 */
func simpleWhereQuery()  {
	// 获取单个记录
	var adminuser model.AdminUser
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((username = 'lisi')) ORDER BY `jh_admin_users`.`id` ASC LIMIT 1
	db.Where("username = ?","lisi").First(&adminuser)
	fmt.Println(adminuser)
	// 获取多个记录
	var adminusers []model.AdminUser
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((status = 1))
	db.Where("status = ?",1).Find(&adminusers)
	fmt.Println(adminusers)
}

/**
	使用结构体、map数据结构、[]int(主键id)方式进行查询
	需要注意一下：当使用结构体的时候，自然是用结构体的属性来设置属性值。但是当使用map数据结构的时候，那就需要用到定义在数据表的表字段来进行设置属性值了。
 */
func advancedWhereQuery()  {
	// 1.使用结构体来构建的where查询条件
	var adminusers []model.AdminUser
	whereStruct := model.AdminUser{
		Status:1,
		LoginFails:2,
	}
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((`jh_admin_users`.`status` = 1) AND (`jh_admin_users`.`login_fails` = 2))
	db.Where(whereStruct).Find(&adminusers)
	fmt.Println(adminusers)

	// 2.使用map方式构建的where查询条件。下面map结构体中定义的key都是数据表对应的字段名
	var adminusers2 []model.AdminUser
	whereMap := map[string]interface{}{
		"status":1,
		"login_fails":2,
	}
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((`jh_admin_users`.`status` = 1) AND (`jh_admin_users`.`login_fails` = 2))
	db.Where(whereMap).Find(&adminusers2)
	fmt.Println(adminusers2)

	// 3.使用主键ID来构建where查询条件
	var adminusers3 []model.AdminUser
	wherePrimarykey := []int{1,3,5}
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((`jh_admin_users`.`id` IN (1,3,5)))
	db.Where(wherePrimarykey).Find(&adminusers3)
	fmt.Println(adminusers3)
}

/**
	Not条件查询
	在Not()函数中，依然可以接受普通字符串、结构体、map数据结构、[]int(主键id)这几种方式
 */
func notWhereQuery()  {
	// 1.使用普通字符串进行的查询。
	var adminusers []model.AdminUser
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND (NOT (login_fails > 1))
	db.Not("login_fails > ?",1).Find(&adminusers)
	fmt.Println(adminusers)

	// 2.使用结构体进行的查询。
	var adminusers2 []model.AdminUser
	whereStruct := model.AdminUser{
		Status:1,
		LoginFails:2,
	}
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((`jh_admin_users`.`status` <> 1) AND (`jh_admin_users`.`login_fails` <> 2))
	db.Not(whereStruct).Find(&adminusers2)
	fmt.Println(adminusers2)

	// 3.使用map结构
	var adminusers3 []model.AdminUser
	whereMap := map[string]interface{}{
		"status":1,
		"login_fails":2,
	}
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((`jh_admin_users`.`status` <> 1) AND (`jh_admin_users`.`login_fails` <> 2))
	db.Not(whereMap).Find(&adminusers2)
	fmt.Println(adminusers3)

	// 4.使用主键ID来构建where查询条件
	var adminusers4 []model.AdminUser
	wherePrimarykey := []int{1,3,5}
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((`jh_admin_users`.`id` NOT IN (1,3,5)))
	db.Not(wherePrimarykey).Find(&adminusers4)
	fmt.Println(adminusers4)
}

/**
	不依赖于Where()或是Not()方法，而是直接使用Find()函数的第二个参数来进行多类型条件查询
 */
func findWhereQuery()  {
	// 1.使用普通字符串进行的查询。
	var adminusers []model.AdminUser
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((login_fails > 1))
	db.Find(&adminusers,"login_fails > ?",1)
	fmt.Println(adminusers)

	// 2.使用结构体进行的查询。
	var adminusers2 []model.AdminUser
	whereStruct := model.AdminUser{
		Status:1,
		LoginFails:2,
	}
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((`jh_admin_users`.`status` = 1) AND (`jh_admin_users`.`login_fails` = 2))
	db.Find(&adminusers2,whereStruct)
	fmt.Println(adminusers2)

	// 3.使用map结构
	var adminusers3 []model.AdminUser
	whereMap := map[string]interface{}{
		"status":1,
		"login_fails":2,
	}
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((`jh_admin_users`.`status` = 1) AND (`jh_admin_users`.`login_fails` = 2))
	db.Find(&adminusers2,whereMap)
	fmt.Println(adminusers3)

	// 4.使用主键ID来构建where查询条件
	var adminusers4 []model.AdminUser
	wherePrimarykey := []int{1,3,5}
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((`jh_admin_users`.`id` IN (1,3,5)))
	db.Find(&adminusers4,wherePrimarykey)
	fmt.Println(adminusers4)

	// 5.查询单个记录
	var adminuser model.AdminUser
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((`jh_admin_users`.`id` = 2))
	db.Find(&adminuser,2)
	fmt.Println(adminuser)
}

/**
	Or条件查询
	使用Or()函数进行查询，在该函数中依然可以接受普通字符串、结构体、map数据结构、[]int(主键id)这几种方式
	从下面的的这个用法示例中，我们可以发现Where()、Or()、Not()这些函数是可以配合使用的，使用查询链的方式来进行组合。
 */
func orWhereQuery()  {
	// 1.使用普通字符串进行的查询。
	var adminusers []model.AdminUser
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((status = 1) OR (login_fails > 1))
	db.Where("status = ?",1).Or("login_fails > ?",1).Find(&adminusers)
	fmt.Println(adminusers)

	// 2.使用结构体进行的查询。
	var adminusers2 []model.AdminUser
	whereStruct := model.AdminUser{
		LoginFails:2,
	}
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((status = 1) OR (`jh_admin_users`.`login_fails` = 2))
	db.Where("status = ?",1).Or(whereStruct).Find(&adminusers2)
	fmt.Println(adminusers2)

	// 3.使用map结构
	var adminusers3 []model.AdminUser
	whereMap := map[string]interface{}{
		"login_fails":2,
	}
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((status = 1) OR (`jh_admin_users`.`login_fails` = 2))
	db.Where("status = ?",1).Or(whereMap).Find(&adminusers2)
	fmt.Println(adminusers3)

	// 4.使用主键ID来构建where查询条件
	var adminusers4 []model.AdminUser
	wherePrimarykey := []int{1,3,5}
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((status = 1) OR (`jh_admin_users`.`id` IN (1,3,5)))
	db.Where("status = ?",1).Or(wherePrimarykey).Find(&adminusers4)
	fmt.Println(adminusers4)
}

/**
	FirstOrInit()和FirstOrCreate()函数。这两个函数所能接收的参数只能是struct或是map
	FirstOrInit()：获取第一个匹配的记录，或者使用给定的条件初始化一个新的记录。如果查询到记录，就会把查询的结果映射到对应的变量上面;如果没有查询到结果，那只会把查询的结果字段映射到对应的变量上面
	FirstOrCreate()：获取第一个匹配的记录。查询到记录，那和FirstOrInit()做的功能一致。如果没有查询到记录，则会先插入一条记录，然后查询当前被插入的记录信息，然后把该记录信息映射到相应的变量上面。
	所以，当使用FirstOrCreate()这个函数来获取一条不存在的记录的时候，是需要执行三条SQL的：①、执行查询。②、执行插入。③、在执行查询(最后根据主键ID来查询)
 */
func firstOrInitQuery()  {
	// 1.用法一：使用结构体方式
	var adminuser model.AdminUser
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((`jh_admin_users`.`username` = 'guolin')) ORDER BY `jh_admin_users`.`id` ASC LIMIT 1
	db.FirstOrInit(&adminuser,model.AdminUser{Username:"guolin"})
	fmt.Println(adminuser)

	// 2.用法二：使用map方式
	var adminuser2 model.AdminUser
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((`jh_admin_users`.`username` = 'guolin')) ORDER BY `jh_admin_users`.`id` ASC LIMIT 1
	// 下面的语句用到了Assign()函数，这个会不管是否查询到结果，都会将adminuser2对象的Status属性值映射为Assign()函数里面设定的值
	db.Assign(model.AdminUser{Status:10}).FirstOrInit(&adminuser2,map[string]interface{}{"username":"guolin"})
	fmt.Println(adminuser2)
	// ---------------------FirstOrCreate()使用---------------------------
	// 1.用法一：使用结构体方式
	var adminuser3 model.AdminUser
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((`jh_admin_users`.`username` = 'guolin')) ORDER BY `jh_admin_users`.`id` ASC LIMIT 1
	db.FirstOrCreate(&adminuser3,model.AdminUser{Username:"guolin"})
	fmt.Println(adminuser3)

	// 2.用法二：使用map方式
	var adminuser4 model.AdminUser
	// SQL查询语句：
	// SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((`jh_admin_users`.`username` = 'guolin')) ORDER BY `jh_admin_users`.`id` ASC LIMIT 1
	// INSERT  INTO `jh_admin_users` (`created_at`,`updated_at`,`deleted_at`,`uuid`,`username`,`status`,`login_fails`) VALUES ('2019-08-27 10:13:06','2019-08-27 10:13:06',NULL,'b3139e4c-5816-405e-a5a3-2e62de6f4684','zhengzexi',10,20)
	// SELECT `age` FROM `jh_admin_users`  WHERE (id = 8)
	// 下面这个语句用到了Attrs()函数，就是当使用FirstOrCreate()函数进行查询的时候，如果没有发现该记录，在新插入的记录的时候，有些字段的值会设置为Attrs()函数设定的值。
	db.Attrs(model.AdminUser{Status:10,LoginFails:20}).FirstOrCreate(&adminuser4,map[string]interface{}{"username":"zhengzexi"})
	fmt.Println(adminuser4)

	// 2.用法二：使用map方式
	var adminuser5 model.AdminUser
	// 下面这个语句同时用到了Attrs()和Assign()函数。其实最终生效的还是Assign()函数定义的值。这就说明，Assign()函数优先级高于Attrs()函数。
	db.Attrs(model.AdminUser{Status:10,LoginFails:20}).Assign(model.AdminUser{Status:5,LoginFails:5}).FirstOrCreate(&adminuser5,map[string]interface{}{"username":"liudekai"})
	fmt.Println(adminuser5)
}

/**
	查询指定字段值.
	Select()函数多接受字符串或是字符串切片格式数据
 */
func selectQuery()  {
	var adminuser model.AdminUser
	// SQL语句：SELECT username,uuid,status,age FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((username = 'lisi')) ORDER BY `jh_admin_users`.`id` ASC LIMIT 1
	db.Select("username,uuid,status,age").Where("username = ?","lisi").First(&adminuser)
	//或是下面这种方式：
	//db.Select([]string{"username","uuid","status","age"}).Where("username = ?","lisi").First(&adminuser)
	fmt.Println(adminuser)
}

/**
	Order()、Limit()、Offset()函数使用
	如果在Order()函数中只是说明了字段，并没有指定何种排序(ASC|DESC)，那默认就是采用的是ASC方式排序
 */

func pageQuery()  {
	var adminusers []model.AdminUser
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((status =1)) ORDER BY `status`,id desc LIMIT 10 OFFSET 0
	db.Where("status =?",1).Order("status").Order("id desc").Limit(10).Offset(0).Find(&adminusers)
	fmt.Println(adminusers)
}

/**
	获取记录总数。使用Count()函数
 */
func countQuery()  {
	var count int
	// SQL查询语句：SELECT count(*) FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL
	db.Model(model.AdminUser{}).Count(&count)
	fmt.Println(count)
}

/**
	指定要查询的数据表。使用Table()函数
	需要注意的是，Table()函数中必须要写完整的表名称，当然就包含了表前缀(如果有的话)和表后缀(如果有的话)
 */
func tableQuery()  {
	var adminuser model.AdminUser
	// SQL查询语句：SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((`jh_admin_users`.`id` = 1))
	db.Table("jh_admin_users").Find(&adminuser,1)
	fmt.Println(adminuser)
}

/**
	Scan()函数的使用
	将查询的结果扫描到另外一个结构体中。下面定义了一个简单的SimpleAdminUser结构体，这个结构体只是包含了model.AdminUser这个结构体的部分字段。
 */
type SimpleAdminUser struct {
	ID int
	Username string
	Age int
	Uuid string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func scanQuery()  {
	var simpleAdminUser []SimpleAdminUser
	// SQL查询语句：SELECT id,username,age,uuid,created_at,updated_at FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((id > 3))
	db.Model(&model.AdminUser{}).Select("id,username,age,uuid,created_at,updated_at").Where("id > ?",3).Scan(&simpleAdminUser)
	fmt.Println(simpleAdminUser)
}

/**
	查询单列字段信息
 */
func pluckQuery()  {
	var status []int
	db.Model(model.AdminUser{}).Where("id > ?",2).Pluck("status",&status)
	fmt.Println(status)
}

/**
	Scopes()函数使用
	简单的说，就是在Scopes()中根据实际需要组合多种查询条件。而这些查询条件都会事先预定好。但是这种条件限定的比较死板(无法接受更多额外业务相关参数)，所以个人觉得这个实际应用场景不会太大
 */
func scopesQuery()  {
	var adminusers []model.AdminUser
	// 最后的SQL查询语句： SELECT * FROM `jh_admin_users`  WHERE `jh_admin_users`.`deleted_at` IS NULL AND ((status = 1) AND (login_fails > 5) AND (age > 25))
	db.Scopes(statusContion,loginFailsContion,ageContion).Find(&adminusers)
	fmt.Println(adminusers)
}
func statusContion(db *gorm.DB) *gorm.DB {
	return db.Where("status = ?",1)
}
func loginFailsContion(db *gorm.DB) *gorm.DB{
	return db.Where("login_fails > ?",5)
}
func ageContion(db *gorm.DB) *gorm.DB{
	return db.Where("age > ?",25)
}
/**
	TODO:Group()、Having()的使用
 */
func groupHavingQuery()  {

}

/**
	TODO:Join()使用
 */

func joinQuery()  {

}

/**
	TODO:Preload()使用
 */
func preloadQuery()  {

}