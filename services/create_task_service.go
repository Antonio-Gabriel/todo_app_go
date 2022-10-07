package services

import (
	"main.go/interfaces"
	"main.go/schemas"
)

type TaskRepositories struct {
	repository interfaces.ITaskRepository
}

func NewTasksRepositories(taskRepository interfaces.ITaskRepository) *TaskRepositories {
	return &TaskRepositories{
		repository: taskRepository,
	}
}

func (tr TaskRepositories) ShowTasks() []schemas.Tasks {
	return tr.repository.Get()
}

func (tr TaskRepositories) ShowTaskById(task_id int) schemas.Tasks {
	return tr.repository.GetById(task_id)
}

func (tr TaskRepositories) CreateNewTask(task schemas.Tasks) (schemas.Tasks, error) {
	res, err := tr.repository.Create(task)
	return *res, err
}

func (tr TaskRepositories) UpdateTask(task schemas.Tasks) (schemas.Tasks, error) {
	res, err := tr.repository.Update(task)
	return *res, err
}

func (tr TaskRepositories) DeleteTask(task_id int) string {
	res := tr.repository.Delete(task_id)
	return res
}
