package utils

import (
	"main.go/schemas"
	"main.go/services"
)

func NewTasksRepositoriesInstance() *services.TaskRepositories {
	return services.NewTasksRepositories(schemas.Tasks{})
}
