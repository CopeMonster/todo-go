package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-go/auth"
	"todo-go/models"
	"todo-go/todo/usecase"
)

func TestGoTodo(t *testing.T) {
	testUser := &models.User{
		Username: "testuser",
		Password: "testpass",
	}

	router := gin.Default()
	group := router.Group("/api", func(ctx *gin.Context) {
		ctx.Set(auth.CtxUserKey, testUser)
	})

	useCase := new(usecase.TodoUseCaseMock)

	RegisterHTTPEndpoints(group, useCase)

	td := &models.Todo{
		ID:          "id",
		Title:       "title",
		Description: "description",
	}

	useCase.On("GetTodo", testUser, "id").Return(td, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/todos/%s", "id"), nil)
	router.ServeHTTP(w, req)

	expectedOut := &getTodoResponse{
		Todo: toTodo(td),
	}

	expectedOutBody, err := json.Marshal(expectedOut)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expectedOutBody), w.Body.String())
}

func TestGoTodos(t *testing.T) {
	testUser := &models.User{
		Username: "testuser",
		Password: "testpass",
	}

	router := gin.Default()
	group := router.Group("/api", func(ctx *gin.Context) {
		ctx.Set(auth.CtxUserKey, testUser)
	})

	useCase := new(usecase.TodoUseCaseMock)

	RegisterHTTPEndpoints(group, useCase)

	tds := make([]*models.Todo, 5)

	for i := 0; i < 5; i++ {
		tds[i] = &models.Todo{
			ID:          "id",
			Title:       "title",
			Description: "description",
		}
	}

	useCase.On("GetTodos", testUser).Return(tds, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/todos", nil)
	router.ServeHTTP(w, req)

	expectedOut := &getTodosResponse{
		Todos: toTodos(tds),
	}

	expectedOutBody, err := json.Marshal(expectedOut)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expectedOutBody), w.Body.String())
}

func TestCreateTodo(t *testing.T) {
	testUser := &models.User{
		Username: "testuser",
		Password: "testpass",
	}

	router := gin.Default()
	group := router.Group("/api", func(ctx *gin.Context) {
		ctx.Set(auth.CtxUserKey, testUser)
	})

	useCase := new(usecase.TodoUseCaseMock)

	RegisterHTTPEndpoints(group, useCase)

	inp := &createTodoInput{
		Title:       "testtitle",
		Description: "testdesc",
	}

	body, err := json.Marshal(inp)
	assert.NoError(t, err)

	useCase.On("CreateTodo", testUser, inp.Title, inp.Description).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/todos", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteTodo(t *testing.T) {
	testUser := &models.User{
		Username: "testuser",
		Password: "testpass",
	}

	router := gin.Default()
	group := router.Group("/api", func(ctx *gin.Context) {
		ctx.Set(auth.CtxUserKey, testUser)
	})

	useCase := new(usecase.TodoUseCaseMock)

	RegisterHTTPEndpoints(group, useCase)

	tdID := "id"

	body, err := json.Marshal(tdID)
	assert.NoError(t, err)

	useCase.On("DeleteTodo", testUser, tdID).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/todos/%s", tdID), bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
