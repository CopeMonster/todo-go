# Todo-Go

## Description
*Todo-Go* is a backend service built with Go for managing tasks. This project aims to provide a fast, reliable, and scalable API for task management applications.

## Installation
To install *Todo-Go*, you need to have Go installed on your machine. Once you have Go installed, you can clone this repository and build the project:

```bash
git clone https://github.com/CopeMonster/todo-go.git
cd todo-go
go build
```

## Usage
After building the project, you can run the API server with:

```bash
go run ./cmd/api/main.go
```
The server will start on the default port 8080. You can then send HTTP requests to the server to interact with the API.

## API Endpoints

### Auth
* `POST /auth/sign-up` Use this endpoint to register a new user to the application.
* `POST /auth/sign-in` Use this endpoint to authenticate a user and log them into the application

### Todos
* `GET /todos/` Use this endpoint to retrieve all tasks from the user's task list
* `GET /todos/:{id}` Use this endpoint to retrieve task from the user's task list by its id
* `POST /todos/` Use this endpoint to add a new task to the user's task list.
* `PUT /todos/:{id}` Use this endpoint to update the details of a specific task from the user's task list. Replace `{id}` with the actual ID of the task.
* `DELETE /todos/:{id}` Use this endpoint to remove a specific task from the user's task list. Replace `{id}` with the actual ID of the task.

## Database
This project uses MongoDB as its primary database. There are 2 collections, `user_collection` and `todo_collection`

## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.

## License
This project is licensed under the [MIT License](https://github.com/CopeMonster/todo-go/blob/master/LICENSE).
