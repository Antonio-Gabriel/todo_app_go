package main

import (
	"fmt"
	"log"
	"net/http"

	"main.go/controllers"
)

func main() {

	// owner := schemas.Owner{
	// 	Name:  "António Gabriel",
	// 	Email: "antoniogabriel@gmail.com",
	// }

	// task := schemas.Tasks{
	// 	TaskName:   "Build a Go lang project",
	// 	Descripion: "The project is a simple task manager or TODO",
	// 	Owner:      owner,
	// 	Done:       false,
	// 	Created_at: time.Now(),
	// }

	// taskRepository := utils.NewTasksRepositoriesInstance()
	// taskRepository.CreateNewTask(task)
	// taskRepository.CreateNewTask(task)
	// taskRepository.CreateNewTask(task)

	http.HandleFunc("/", controllers.InitHandler)
	http.HandleFunc("/tasks", controllers.ListTasksController)
	http.HandleFunc("/tasks/create", controllers.CreateTasksController)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}

	// owner := schemas.Owner{
	// 	Name:  "António Gabriel",
	// 	Email: "antoniogabriel@gmail.com",
	// }

	// task := schemas.Tasks{
	// 	TaskName:   "Build a Go lang project",
	// 	Descripion: "The project is a simple task manager or TODO",
	// 	Owner:      owner,
	// 	Done:       false,
	// 	Created_at: time.Now(),
	// }

	// taskRepository := utils.NewTasksRepositoriesInstance()
	// taskRepository.CreateNewTask(task)
	// taskRepository.CreateNewTask(task)
	// taskRepository.CreateNewTask(task)

	// task_to_update := schemas.Tasks{
	// 	Id:         1,
	// 	TaskName:   "Alter to other text",
	// 	Descripion: "The project is a simple task manager or TODO",
	// 	Owner:      owner,
	// 	Done:       false,
	// 	Created_at: time.Now(),
	// }

	// taskRepository := utils.NewTasksRepositoriesInstance()

	// // fmt.Println("\n")

	// taskRepository.CreateNewTask(task)
	// taskRepository.CreateNewTask(task)
	// taskRepository.CreateNewTask(task)

	// // fmt.Println("\n")

	// taskRepository.UpdateTask(task_to_update)

	// taskJSON, _ := json.Marshal(taskRepository.ShowTasks())

	// fmt.Println(string(taskJSON))

	// fmt.Println(taskRepository.ShowTasks())
}
