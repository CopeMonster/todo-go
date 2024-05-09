package localstorage

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"todo-go/models"
	"todo-go/todo"
)

func TestGetTodos(t *testing.T) {
	userID := "id"

	user := &models.User{
		ID: userID,
	}

	storage := NewTodoLocalStorage()

	for i := 0; i < 10; i++ {
		td := &models.Todo{
			ID:     fmt.Sprintf("id%d", i),
			UserID: userID,
		}

		err := storage.CreateTodo(context.Background(), user, td)
		assert.NoError(t, err)
	}

	returnedTds, err := storage.GetTodos(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, 10, len(returnedTds))
}

func TestGetTodo(t *testing.T) {
	userID := "id"

	user := &models.User{
		ID: userID,
	}

	storage := NewTodoLocalStorage()

	tdID := "id10"

	td := &models.Todo{
		ID:     tdID,
		UserID: userID,
	}

	err := storage.CreateTodo(context.Background(), user, td)
	assert.NoError(t, err)

	returnedTd, err := storage.GetTodo(context.Background(), user, tdID)
	assert.NoError(t, err)
	assert.Equal(t, td, returnedTd)
}

func TestUpdateTodo(t *testing.T) {
	userID := "id"

	user := &models.User{
		ID: userID,
	}

	storage := NewTodoLocalStorage()

	tdID := "id10"

	td := &models.Todo{
		ID:          tdID,
		UserID:      userID,
		Title:       "Create tests",
		Description: "Just test description",
	}

	err := storage.CreateTodo(context.Background(), user, td)
	assert.NoError(t, err)

	updatedTd := &models.Todo{
		Title:       "Create tests",
		Description: "Just test description 2",
		Done:        true,
	}

	err = storage.UpdateTodo(context.Background(), user, tdID, updatedTd)
	assert.NoError(t, err)
}

func TestDeleteTodo(t *testing.T) {
	userID1 := "id1"
	userID2 := "id2"

	user1 := &models.User{
		ID: userID1,
	}

	user2 := &models.User{
		ID: userID2,
	}

	tdID := "tdID"
	td := &models.Todo{
		ID:     tdID,
		UserID: userID1,
	}

	storage := NewTodoLocalStorage()

	err := storage.CreateTodo(context.Background(), user1, td)
	assert.NoError(t, err)

	err = storage.DeleteTodo(context.Background(), user1, tdID)
	assert.NoError(t, err)

	err = storage.CreateTodo(context.Background(), user1, td)
	assert.NoError(t, err)

	err = storage.DeleteTodo(context.Background(), user2, tdID)
	assert.Error(t, err)
	assert.Equal(t, err, todo.ErrTodoNotFound)
}
