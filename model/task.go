package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	//外键关联到user表中
	//反引号表示：gorm的标签，gorm会根据标签内容来创建
	User      User   `gorm:"ForeignKey:Uid"`
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index;not null"`
	Status    int    `gorm:"default:'0'"` //0表示未完成，1表示已完成
	Content   string `gorm:"type:longtext"`
	StartTime int64  //备忘录开始时间
	EndTime   int64  //备忘录结束时间
}
