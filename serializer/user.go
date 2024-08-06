package serializer

import "todo_list/model"

// service中构建的为数据库中的结构体通过serializer进行序列化
type User struct {
	ID       uint   `json:"id" form:"id" example:"1"` //用户ID
	UserName string `json:"user_name" form="username" example:"xhq"`
	Status   string `json:"status" form:"status"`       //用户状态
	CreateAt int64  `json:"create_at" form:""create_at` //创建时间
}

func BuildUser(user model.User) User {
	//接收数据库表中数据结构的user，创建一个serial的user结构体
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
}
