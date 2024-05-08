package auth

import (
	"context"
	"todo-go/models"
)

const CtxUserKey = "user"

type UseCase interface {
	SignUp(ctx context.Context, username string, password string) error
	SignIn(ctx context.Context, username string, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*models.User, error)
}
