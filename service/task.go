package service

import (
	"time"
	"todo_list/model"
	"todo_list/serializer"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0是未做 1是已做
}

type ShowTaskService struct {
}

type ListTaskService struct {
	PageSize int `json:"page_size" form:"page_size"`
	PageNum  int `json:"page_num" form:"page_num"`
}

type UpdateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0是未做 1是已做
}

type SearchTaskService struct {
	Info     string `json:"info" form:"info"`
	PageSize int    `json:"page_size" form:"page_size"`
	PageNum  int    `json:"page_num" form:"page_num"`
}

type DeleteTaskService struct {
	Tid uint `json:"tid" form:"tid"`
}

// 定义结构体的方法

// 创建一条备忘录
func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	code := 200
	//first()：在表中找到的第一条数据赋值给user
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Content:   service.Content,
		Status:    0,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "创建备忘录失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "创建成功",
	}
}

// 查询展示某一条备忘录
func (service *ShowTaskService) Show(tid string) serializer.Response {
	var task model.Task
	code := 200
	err := model.DB.First(&task, tid).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "查询失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTask(task),
	}
}

// 列表展示用户所有的备忘录
func (service *ListTaskService) List(uid uint) serializer.Response {
	//切片类型
	var task []model.Task
	count := 0
	//默认每一页10条备忘录数据
	if service.PageSize == 0 {
		service.PageSize = 10
	}
	//多表查询
	//1、model.DB.Model:使用GORM库开始一个新的数据库查询，指定查询的模型
	//2、.Preload：预加载与task相关的user关联对象，意味着查询将同时获取任务即其关联的用户信息
	//3、count：计算查询结果输了
	//4、service.PageNum是当前请求的页码，页码从1开始，需要-1来适应数据库的0索引
	//5、find：执行查询，并将结果填充到task切片中
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Count(&count).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&task)
	return serializer.Response{
		Status: 200,
		//
		Data: serializer.BuildListResponse(serializer.BuildTasks(task), uint(count)),
		// Data:   (serializer.BuildTasks(task)),          不想用这种格式返回，这样的方式看不到备忘录总数
	}
}

func (service *UpdateTaskService) Update(tid string) serializer.Response {
	var task model.Task
	model.DB.First(&task, tid)
	task.Content = service.Content
	task.Title = service.Title
	task.Status = service.Status
	//gorm的save方法用于更新数据库的数据,在这个过程需要进行错误处理
	updateResult := model.DB.Save(&task)
	if updateResult.Error != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "更新数据库发生错误" + updateResult.Error.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		//不仅要返回成功，也需要返回查询得到的数据
		Data: serializer.BuildTask(task),
		Msg:  "更新完成！",
	}
}

// 查询某个用户的备忘录（模糊查询）
func (service *SearchTaskService) Search(uid uint) serializer.Response {
	var task []model.Task
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 10
	}
	//我需要得到某个用户的所有数据，所以最终获得的是task数据，但需要预加载user表
	model.DB.Model(&model.Task{}).Preload("User").
		Where("uid=?", uid).Where("title LIKE ? OR content LIKE ?", "%"+service.Info+"%", "%"+service.Info+"%").
		Count(&count).Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&task)

	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildListResponse(serializer.BuildTasks(task), uint(count)),
		Msg:    "模糊查询成功！",
	}

}

func (service *DeleteTaskService) Delete(tid string) serializer.Response {
	var task model.Task
	model.DB.First(&task, tid)
	if err := model.DB.Delete(task).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库删除错误" + err.Error(),
		}
	} else {
		return serializer.Response{
			Status: 200,
			Msg:    "删除成功",
		}
	}
}
