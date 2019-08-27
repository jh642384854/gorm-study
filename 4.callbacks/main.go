package main

/**
	这个应用主要在模型定义上面，可以为一个模型定义多个方法，在CURD的不同操作会被调用

	1.创建对象时
	Creating an object，创建对象时可用的 hooks

	// 开始事务
	BeforeSave
	BeforeCreate
	// 在关联前保存
	// 更新时间戳 `CreatedAt`, `UpdatedAt`
	// save self
	// 重新加载具有默认值的字段，其值为空
	// 在关联后保存
	AfterCreate
	AfterSave
	// 提交或回滚事务

	2.更新对象时
	Updating an object，更新对象时可用的 hooks

	// begin transaction 开始事物
	BeforeSave
	BeforeUpdate
	// save before associations 保存前关联
	// update timestamp `UpdatedAt` 更新 `UpdatedAt` 时间戳
	// save self 保存自己
	// save after associations 保存后关联
	AfterUpdate
	AfterSave
	// commit or rollback transaction 提交或回滚事务

	3.删除对象时
	Deleting an object，删除对象时可用的 hooks

	// begin transaction 开始事务
	BeforeDelete
	// delete self 删除自己
	AfterDelete
	// commit or rollback transaction 提交或回滚事务

	4.查询对象时
	Querying an object，查询对象时可用的 hooks

	// load data from database 从数据库加载数据
	// Preloading (eager loading) 预加载（加载）
	AfterFind
 */
func main() {

}
