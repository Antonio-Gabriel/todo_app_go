package controllers

import (
	"encoding/json"
	"net/http"

	"main.go/schemas"
	"main.go/utils"
)

func CreateTasksController(res http.ResponseWriter, req *http.Request) {

	utils.HttpMethod("POST", res, req)

	var task schemas.Tasks

	payload := req.Body

	defer req.Body.Close()

	err := json.NewDecoder(payload).Decode(&task)
	if err != nil {
		HandlerMessage := []byte(`
			{
				"success": false,
				"message": "Error parsing the task data"
			}
		`)

		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
	}

	taskService := utils.NewTasksRepositoriesInstance()

	taskSaved, _ := taskService.CreateNewTask(task)

	taskJSON, _ := json.Marshal(taskSaved)

	utils.ReturnJsonResponse(res, http.StatusOK, taskJSON)
}
