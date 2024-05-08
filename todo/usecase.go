package todo

import (
	"context"
	"todo-go/models"
)

type UseCase interface {
	GetTodo(ctx context.Context, user *models.User, id string) (*models.Todo, error)
	GetTodos(ctx context.Context, user *models.User) ([]*models.Todo, error)
	CreateTodo(ctx context.Context, user *models.User, title string, description string) error
	UpdateTodo(ctx context.Context, user *models.User, id string, title string, description string, done bool) error
	DeleteTodo(ctx context.Context, user *models.User, id string) error
}
