package api

import (
	"todo_list/pkg/utils"
	"todo_list/service"

	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

// 新增备忘录
func CreateTask(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	// 创建一个服务
	var createTask service.CreateTaskService
	//err==nil  而不是err!=nil
	if err := c.ShouldBind(&createTask); err != nil {
		logging.Error(err)
		//绑定数据时存在错误，在api/main.go中定义序列化器，让其返回json格式
		c.JSON(400, ErrorResponse(err))
		//return防止进一步执行
		return
	}
	//若Title为必要字段，则添加判断逻辑
	if createTask.Title == "" {
		//c.json：gin框架用来返回json格式响应的函数，包括：状态码、内容
		c.JSON(400, gin.H{"error": "Title is required"})
		return
	}
	res := createTask.Create(claim.Id)
	c.JSON(200, res)
}

// 展示备忘录
func ShowTask(c *gin.Context) {
	// 拿到id，直接去数据库中查询，直接通过gorm遍历数据库中的数据
	// claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	// 创建一个服务
	var showTask service.ShowTaskService
	//err==nil  而不是err!=nil
	if err := c.ShouldBind(&showTask); err != nil {
		logging.Error(err)
		//绑定数据时存在错误，在api/main.go中定义序列化器，让其返回json格式
		c.JSON(400, ErrorResponse(err))
		//return防止进一步执行
		return
	}
	//若Title为必要字段，则添加判断逻辑

	//传递：用户上传的任务id（get请求中的参数） 唯一标识
	res := showTask.Show(c.Param("id"))
	c.JSON(200, res)
}

// 查询某用户的所有的备忘录
func ListTask(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	var listTask service.ListTaskService

	if err := c.ShouldBind(&listTask); err != nil {
		logging.Error(err)
		//绑定数据时存在错误，在api/main.go中定义序列化器，让其返回json格式
		c.JSON(400, ErrorResponse(err))
		//return防止进一步执行
		return
	}
	res := listTask.List(claim.Id)
	c.JSON(200, res)
}

// 更新一条备忘录（需要用户鉴权、tid）
func UpdateTask(c *gin.Context) {
	// claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	var updateTask service.UpdateTaskService
	if err := c.ShouldBind(&updateTask); err != nil {
		logging.Error(err)
		//绑定数据时存在错误，在api/main.go中定义序列化器，让其返回json格式
		c.JSON(400, ErrorResponse(err))
		//return防止进一步执行
		return
	}
	//此处c.param()返回的为string类型数据
	res := updateTask.Update(c.Param("id"))
	c.JSON(200, res)
}

// 模糊查询用户的数据
func SearchTask(c *gin.Context) {
	var searchTask service.SearchTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))

	if err := c.ShouldBind(&searchTask); err != nil {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
		return
	}

	res := searchTask.Search(claim.Id)
	c.JSON(200, res)
}

// 删除用户的某条数据
func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService

	if err := c.ShouldBind(&deleteTask); err != nil {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
		return
	}

	res := deleteTask.Delete(c.Param("id"))
	c.JSON(200, res)
}
