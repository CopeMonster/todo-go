package http

import (
	"github.com/gin-gonic/gin"
	"todo-go/auth"
)

func RegisterHTTPEndpoints(router *gin.Engine, useCase auth.UseCase) {
	handler := NewHandler(useCase)

	authEndpoints := router.Group("/auth")

	{
		authEndpoints.POST("/sign-up", handler.SignUp)
		authEndpoints.POST("/sign-in", handler.SignIn)
	}
}
