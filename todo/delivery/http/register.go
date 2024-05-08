package http

import (
	"github.com/gin-gonic/gin"
	"todo-go/todo"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, useCase todo.UseCase) {
	handler := NewHandler(useCase)

	todos := router.Group("/todos")

	{
		todos.GET("", handler.GetTodos)
		todos.GET("/:id", handler.GetTodo)
		todos.POST("", handler.CreateTodo)
		todos.PUT("/:id", handler.UpdateTodo)
		todos.DELETE("/:id", handler.DeleteTodo)
	}
}
