package localstorage

import (
	"context"
	"sync"
	"todo-go/models"
	"todo-go/todo"
)

type TodoLocalStorage struct {
	todos map[string]*models.Todo
	mutex *sync.Mutex
}

func NewTodoLocalStorage() *TodoLocalStorage {
	return &TodoLocalStorage{
		todos: make(map[string]*models.Todo),
		mutex: new(sync.Mutex),
	}
}

func (s *TodoLocalStorage) GetTodo(ctx context.Context, user *models.User, id string) (*models.Todo, error) {
	for _, td := range s.todos {
		if td.UserID == user.ID {
			return td, nil
		}
	}

	return nil, nil
}

func (s *TodoLocalStorage) GetTodos(ctx context.Context, user *models.User) ([]*models.Todo, error) {
	todos := make([]*models.Todo, 0)

	s.mutex.Lock()

	for _, td := range s.todos {
		if td.UserID == user.ID {
			todos = append(todos, td)
		}
	}

	s.mutex.Unlock()

	return todos, nil
}

func (s *TodoLocalStorage) CreateTodo(ctx context.Context, user *models.User, td *models.Todo) error {
	td.UserID = user.ID

	s.mutex.Lock()
	s.todos[td.ID] = td
	s.mutex.Unlock()

	return nil
}

func (s *TodoLocalStorage) UpdateTodo(ctx context.Context, user *models.User, id string, td *models.Todo) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.todos[id]; !ok {
		return todo.ErrTodoNotFound
	}

	td.UserID = user.ID
	s.todos[id] = td

	return nil
}

func (s *TodoLocalStorage) DeleteTodo(ctx context.Context, user *models.User, id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	td, exist := s.todos[id]

	if exist && td.UserID == user.ID {
		delete(s.todos, id)
		return nil
	}

	return todo.ErrTodoNotFound
}
