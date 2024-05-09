package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo-go/auth/repository/mock"
	"todo-go/models"
)

func TestAuthFlow(t *testing.T) {
	repository := new(mock.UserStorageMock)
	useCase := NewAuthUseCase(repository, "salt", []byte("secret"), 86400)

	var (
		username = "user"
		password = "pass"

		ctx = context.Background()

		user = &models.User{
			Username:    username,
			Password:    "11f5639f22525155cb0b43573ee4212838c78d87",
			CreatedTime: time.Now(),
		}
	)

	repository.On("CreateUser", user).Return(nil)

	err := useCase.SignUp(ctx, username, password)
	assert.NoError(t, err)

	repository.On("GetUser", user.Username, user.Password).Return(user, nil)
	token, err := useCase.SignIn(ctx, username, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedUser, err := useCase.ParseToken(ctx, token)

	assert.NoError(t, err)
	assert.Equal(t, user.ID, parsedUser.ID)
	assert.Equal(t, user.Username, parsedUser.Username)
	assert.Equal(t, user.Password, parsedUser.Password)
	assert.Equal(t, user.CreatedTime.Unix(), parsedUser.CreatedTime.Unix())
}
