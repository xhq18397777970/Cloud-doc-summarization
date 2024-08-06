// 将数据结构做迁移，映射到数据库中
package model

func migration() {
	//自动迁移模式
	//由于DB在model的package下，不需要导入model包，直接调用
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}).
		AutoMigrate(&Task{})
		//设置task的外键，关联到user表
	DB.Model(&Task{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")
}
