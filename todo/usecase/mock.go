package usecase

import (
	"context"
	"github.com/stretchr/testify/mock"
	"todo-go/models"
)

type TodoUseCaseMock struct {
	mock.Mock
}

func (m *TodoUseCaseMock) GetTodo(ctx context.Context, user *models.User, id string) (*models.Todo, error) {
	args := m.Called(user, id)

	return args.Get(0).(*models.Todo), args.Error(1)
}

func (m *TodoUseCaseMock) GetTodos(ctx context.Context, user *models.User) ([]*models.Todo, error) {
	args := m.Called(user)

	return args.Get(0).([]*models.Todo), args.Error(1)
}

func (m *TodoUseCaseMock) CreateTodo(ctx context.Context, user *models.User, title string, description string) error {
	args := m.Called(user, title, description)

	return args.Error(0)
}

func (m *TodoUseCaseMock) UpdateTodo(ctx context.Context, user *models.User, id string, title string, description string, done bool) error {
	args := m.Called(user, id, title, description, done)

	return args.Error(0)
}

func (m *TodoUseCaseMock) DeleteTodo(ctx context.Context, user *models.User, id string) error {
	args := m.Called(user, id)

	return args.Error(0)
}
