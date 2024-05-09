package localstorage

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo-go/auth"
	"todo-go/models"
)

func TestGetUser(t *testing.T) {
	storage := NewUserLocalStorage()

	id1 := "id"

	user := &models.User{
		ID:          id1,
		Username:    "user",
		Password:    "password",
		CreatedTime: time.Now(),
	}

	err := storage.CreateUser(context.Background(), user)
	assert.NoError(t, err)

	returnedUser, err := storage.GetUser(context.Background(), "user", "password")
	assert.NoError(t, err)
	assert.Equal(t, user, returnedUser)

	returnedUser, err = storage.GetUser(context.Background(), "user", "wrong_password")
	assert.Error(t, err)
	assert.Equal(t, err, auth.ErrUserNotFound)
}
