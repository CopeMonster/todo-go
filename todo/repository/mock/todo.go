package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"todo-go/models"
)

type TodoStorageMock struct {
	mock.Mock
}

func (s *TodoStorageMock) GetTodo(ctx context.Context, user *models.User, id string) (*models.Todo, error) {
	args := s.Called(user, id)

	return args.Get(0).(*models.Todo), args.Error(1)
}

func (s *TodoStorageMock) GetTodos(ctx context.Context, user *models.User) ([]*models.Todo, error) {
	args := s.Called(user)

	return args.Get(0).([]*models.Todo), args.Error(1)
}

func (s *TodoStorageMock) CreateTodo(ctx context.Context, user *models.User, td *models.Todo) error {
	args := s.Called(user, td)

	return args.Error(0)
}

func (s *TodoStorageMock) UpdateTodo(ctx context.Context, user *models.User, id string, td *models.Todo) error {
	args := s.Called(user, id, td)

	return args.Error(0)
}

func (s *TodoStorageMock) DeleteTodo(ctx context.Context, user *models.User, id string) error {
	args := s.Called(user, id)

	return args.Error(0)
}
