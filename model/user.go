package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	//gorm建表会默认提供三个变量：CreatedAt time.Time、time.Time、DeletedAt，并支持软删除（删除数据只是记录删除信息，数据仍然在数据库中）
	gorm.Model
	//大写，便于gorm找到该全局变量
	UserName string
	//存储密文
	PasswordDigest string
}

// 加密
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// 解密(在登陆时的操作) 验证密码
func (user *User) CheckPassword(password string) bool {
	//将用户上传的明文密码，与数据库中存储的密文密码进行比对
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
