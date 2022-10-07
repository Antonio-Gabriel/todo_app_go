# TODO LIST

A simple project using Go and native resources for build this application and explore a little over data structure in Go. I really liked it though don't have a many resources or big resources to work with data structure such as python, javascript etc.

But it's a good stack, to offer a good experience and performance of develop, I think that yh. That's Good.

## Tasks that the app already make

List of tasks that the app already make:

Some tasks already make at services module and only don't have a controller that apoint to some endpoint, but exists on service module:

- [x] Create a task
- [x] List all tasks
- [ ] List task by id
- [ ] Marke as done
- [ ] Update task
- [ ] Delete task

## Endpoints availables

    - POST(Create task): [
        Url: http://localhost:8888/tasks/create,
        payload: {
            {
                "TaskName": "Build a basic api",
                "Descripion": "The project is a simple task manager or TODO",
                "Owner": {
                    "Name": "António Gabriel",
                    "Email": "antoniogabriel@gmail.com"
                },
                "Done": true,
                "Created_at": "2022-10-07T00:24:28.980192046+01:00"
            }
        }
    ]

    - GET(All tasks): [
        Url: http://localhost:8888/tasks
    ]

In the moment only this routes that runs.

# License

Released in 2022. This project is under the [MIT license](LICENSE)

Made by [António Gabriel](https://github.com/Antonio-Gabriel)