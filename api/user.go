package api

import (
	"todo_list/service"

	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

// 该api用于接收上下文数据，并传递给service
func UserRegister(c *gin.Context) {
	//创建一个服务
	var userRegister service.UserService
	//绑定，相当于将上下文传输的数据传递给该结构体

	if err := c.ShouldBind(&userRegister); err != nil {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
		//打印一下传输过来，并赋值的数据
		return
	}
	//若用户名或密码为空则返回错误
	if userRegister.UserName == "" || userRegister.Password == "" {
		c.JSON(400, gin.H{"error": "user_name & password is required"})
		return
	}
	res := userRegister.Register()
	c.JSON(200, res)

}
func UserLogin(c *gin.Context) {
	//创建一个服务
	var userLogin service.UserService
	//对它进行绑定
	//前段发送的字段名称，要与UserService结构体中的字段标签的名称匹配
	//err==nil    而不是err	!=nil
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}
