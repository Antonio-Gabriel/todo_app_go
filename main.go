package main

import (
	"fmt"
	"log"
	"net/http"

	"main.go/controllers"
)

func main() {

	http.HandleFunc("/", controllers.InitHandler)
	http.HandleFunc("/tasks", controllers.ListTasksController)
	http.HandleFunc("/tasks/create", controllers.CreateTasksController)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
