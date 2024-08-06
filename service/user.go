package service

import (
	"fmt"
	"todo_list/model"
	"todo_list/pkg/e"
	"todo_list/pkg/utils"
	"todo_list/serializer"

	"github.com/jinzhu/gorm"
)

// 接收用户上传的2个参数
type UserService struct {
	//要求用户名最小3位，最大15位
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

// 注册的逻辑
// 定义了属于UserService结构体实例化对象的方法，该对象的方法
func (service *UserService) Register() serializer.Response { //返回一个json格式
	//创建一个user的结构体，若没有该用户，则向数据库中插入
	//验证一下数据库中有没有这个用户
	var user model.User
	//用于接收，是否存在该用户
	var count int
	//1、model.DB为GORM的数据库实例，用于执行数据库操作
	//2、Model(&model.User{})：表示要操作的数据库表对应的是User模型
	//3、First(&user)：获取查询结构的第一行数据，并赋值给user变量
	//4、Count(&count)：技术查询结果行数
	fmt.Println(service.UserName)
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).
		First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: e.ErrorExistUser,
			Msg:    "已经有这个人了，不需要再注册了",
		}
	}
	//若没有这个人，存入该数据,并对密码进行加密
	user.UserName = service.UserName
	//将用户传递来的密码进行加密
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: e.InvalidParams,
			//加密过程中发生错误
			Msg: err.Error(),
		}
	}
	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: e.ERROR,
			Msg:    "数据库操作错误",
		}
	}
	return serializer.Response{
		Status: e.SUCCESS,
		Msg:    "用户注册成功",
	}

}

// 登陆逻辑
func (service *UserService) Login() serializer.Response {
	var user model.User
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: e.ErrorNotExistUser,
				Msg:    "用户不存在,请先登录",
			}
		}
		//如果不是用户不存在，而是其他不可抗拒的因素导致的错误
		return serializer.Response{
			Status: e.ERROR,
			Msg:    "[数据库错误]\n" + err.Error(),
		}
	}
	//调用在model下定义的结构体方法
	if user.CheckPassword(service.Password) == false {
		return serializer.Response{
			Status: e.InvalidParams,
			Msg:    "密码错误",
		}
	}

	//当身份验证成功时，发一个token，为了其他功能需要身份验证，前端存储
	//如创建备忘录，这个功能需要token，不然不知道是谁创建的备忘录

	token, err := utils.GenerateToken(user.ID, service.UserName, service.Password)
	if err != nil {
		return serializer.Response{
			Status: e.ERROR,
			Msg:    "token签发错误" + err.Error(),
		}
	}
	return serializer.Response{
		Status: e.SUCCESS,
		//返回一个携带token的用户信息
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
		Msg: "登陆成功",
	}

}
