package schemas

import (
	"errors"
	"time"
)

type Tasks struct {
	Id         int       `json:id`
	TaskName   string    `json: task`
	Descripion string    `json: desc`
	Owner      Owner     `json: owner`
	Done       bool      `json: done`
	Created_at time.Time `json: created_at`
}

var storage = make([]Tasks, 0)

func (tk Tasks) Get() []Tasks {
	return storage
}

func (tk Tasks) GetById(task_id int) Tasks {
	res := filter(storage, func(task Tasks) bool {
		return task.Id == task_id
	})

	return res[0]
}

func (tk Tasks) Create(task Tasks) (*Tasks, error) {
	err := task.validate()

	if err != nil {
		return nil, errors.New(err.Error())
	}

	task.Id = len(storage) + 1

	storage = append(storage, task)

	return &task, nil
}

func (tk Tasks) Update(task Tasks) (*Tasks, error) {
	res := filter(storage, func(taskFilter Tasks) bool {
		return taskFilter.Id == task.Id
	})

	if res != nil {
		for i, d := range storage {
			if res[0].Id == d.Id {
				storage[i] = task
			}
		}
		return &task, nil
	}

	return nil, errors.New("error to update the task")
}

func (tk Tasks) Delete(task_id int) string {
	res := filter(storage, func(task Tasks) bool {
		return task.Id == task_id
	})

	if res != nil {
		for i, d := range storage {
			if res[0].Id == d.Id {
				storage = append(storage[:i], storage[i+1:]...)
			}
		}

		return "The task was deleted successfully"
	}

	return "Please, input a valid task id"
}

func (task Tasks) validate() error {

	if task.TaskName != "" ||
		task.Owner.Name != "" ||
		task.Owner.Email != "" {

		return nil
	}

	return errors.New("error")
}

func filter(data []Tasks, f func(Tasks) bool) []Tasks {
	fltd := make([]Tasks, 0)

	for _, task := range data {
		if f(task) {
			fltd = append(fltd, task)
		}
	}

	return fltd
}
