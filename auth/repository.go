package auth

import (
	"context"
	"todo-go/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, username string, password string) (*models.User, error)
}
