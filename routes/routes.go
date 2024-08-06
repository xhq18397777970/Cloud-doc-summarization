package routes

import (
	"todo_list/api"
	"todo_list/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// gin.Engine：Gin框架的核心
func NewRouter() *gin.Engine {
	r := gin.Default()
	//创建基于cookie的会话存储，需要密钥作为参数，此处用硬编码密钥
	store := cookie.NewStore([]byte("something-very-secret"))
	//使用gin的use方法添加一个中间价，用于处理会话管理，会话名称为mysession，使用的存储为上一步创建的store
	r.Use(sessions.Sessions("mysession", store))
	//创建路由组v1，所有在这个组下的路由都会加上前缀
	v1 := r.Group("api/v1")
	{
		//当该路由被访问，会调用对应的函数处理
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		//定义新的子路由组，由于没有添加额外的路由前缀，继承api/v1
		authed := v1.Group("/")
		//利用中间价进行鉴权，只有通过鉴权的请求才可以访问该组下的路由
		authed.Use(middleware.JWT())
		{
			//authed为已设置好鉴权中间件的路由组
			authed.POST("task", api.CreateTask)
			//:id：常见的RESTful API设计模式，会根据用户上传的替换:id
			authed.GET("task/:id", api.ShowTask)
			//在restful API中，get请求常用于
			authed.GET("tasks", api.ListTask)
			authed.PUT("task/:id", api.UpdateTask)
			authed.POST("search", api.SearchTask)
			authed.DELETE("task/:id", api.DeleteTask)
		}

	}
	return r
}
