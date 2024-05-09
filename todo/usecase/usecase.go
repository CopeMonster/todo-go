package usecase

import (
	"context"
	"time"
	"todo-go/models"
	"todo-go/todo"
)

type TodoUseCase struct {
	todoRepository todo.Repository
}

func NewTodoUseCase(todoRepository todo.Repository) *TodoUseCase {
	return &TodoUseCase{
		todoRepository: todoRepository,
	}
}

func (t TodoUseCase) GetTodo(ctx context.Context, user *models.User, id string) (*models.Todo, error) {
	return t.todoRepository.GetTodo(ctx, user, id)
}

func (t TodoUseCase) GetTodos(ctx context.Context, user *models.User) ([]*models.Todo, error) {
	return t.todoRepository.GetTodos(ctx, user)
}

func (t TodoUseCase) CreateTodo(ctx context.Context, user *models.User, title string, description string) error {
	td := &models.Todo{
		UserID:       user.ID,
		Title:        title,
		Description:  description,
		Done:         false,
		CreatedTime:  time.Now(),
		ModifiedTime: time.Now(),
	}

	return t.todoRepository.CreateTodo(ctx, user, td)
}

func (t TodoUseCase) UpdateTodo(ctx context.Context, user *models.User, id string, title string, description string, done bool) error {
	td := &models.Todo{
		Title:        title,
		Description:  description,
		Done:         done,
		ModifiedTime: time.Now(),
	}

	return t.todoRepository.UpdateTodo(ctx, user, id, td)
}

func (t TodoUseCase) DeleteTodo(ctx context.Context, user *models.User, id string) error {
	return t.todoRepository.DeleteTodo(ctx, user, id)
}
