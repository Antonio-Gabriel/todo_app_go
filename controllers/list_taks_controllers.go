package controllers

import (
	"encoding/json"
	"net/http"

	"main.go/utils"
)

func ListTasksController(res http.ResponseWriter, req *http.Request) {
	utils.HttpMethod("GET", res, req)

	tasksService := utils.NewTasksRepositoriesInstance()
	taskJSON, _ := json.Marshal(tasksService.ShowTasks())

	utils.ReturnJsonResponse(res, http.StatusOK, taskJSON)
}
