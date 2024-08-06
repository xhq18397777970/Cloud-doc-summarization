package serializer

import "todo_list/model"

type Task struct {
	ID        uint   `json:"id" example:"1"`
	Title     string `json:"title" example:"吃饭"`
	Status    int    `json:"status" example:"0"`
	Content   string `json:"content" example:"睡觉"`
	View      uint64 `json:"view" example:"32"`
	StartTime int64  //备忘录开始时间
	EndTime   int64  //备忘录结束时间
	CreatedAt int64
}

func BuildTask(task model.Task) Task {
	return Task{
		ID:        task.ID,
		Title:     task.Title,
		Status:    task.Status,
		Content:   task.Content,
		CreatedAt: task.CreatedAt.Unix(),
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
}

func BuildTasks(items []model.Task) (tasks []Task) {
	for _, item := range items {
		task := BuildTask(item)
		tasks = append(tasks, task)
	}
	return tasks
}
