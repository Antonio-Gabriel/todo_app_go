package interfaces

import (
	"main.go/schemas"
)

type ITaskRepository interface {
	Get() []schemas.Tasks
	GetById(task_id int) schemas.Tasks
	Create(task schemas.Tasks) (*schemas.Tasks, error)
	Update(task schemas.Tasks) (*schemas.Tasks, error)
	Delete(task_id int) string
}
