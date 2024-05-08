package todo

import (
	"context"
	"todo-go/models"
)

type Repository interface {
	GetTodo(ctx context.Context, user *models.User, id string) (*models.Todo, error)
	GetTodos(ctx context.Context, user *models.User) ([]*models.Todo, error)
	CreateTodo(ctx context.Context, user *models.User, td *models.Todo) error
	UpdateTodo(ctx context.Context, user *models.User, id string, td *models.Todo) error
	DeleteTodo(ctx context.Context, user *models.User, id string) error
}
