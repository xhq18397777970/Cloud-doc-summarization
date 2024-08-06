package model

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 返回DB为全局变量，供其他模块，对数据库进行操作
var DB *gorm.DB

func DataBase(connstring string) {

	// 接收一个path，进行数据库连接
	db, err := gorm.Open("mysql", connstring)
	if err != nil {
		fmt.Println(err)
		panic("mysql数据库连接错误")
	}
	fmt.Println("数据库连接成功")
	//gorm可以输出日志文件
	db.LogMode(true)
	//若gin的版本为运行环境则不输出日志文件
	if gin.Mode() == "release" {
		db.LogMode(false)
	}

	//默认为false，gorm建表时，默认自动加上s，如user变为users。此处设置true
	db.SingularTable(true)
	//设置连接池
	db.DB().SetMaxIdleConns(20)
	//设置最大连接数
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	//gorm model连接数据库的方式，用gorm.open方式连接，此时db即为数据库操作对象
	DB = db

	//将定义好的结构体，通过migration，直接迁移为数据库中的一张表
	migration()
}
