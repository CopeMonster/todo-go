package http

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-go/auth"
	"todo-go/auth/usecase"
	"todo-go/models"
)

func TestAuthMiddleware(t *testing.T) {
	router := gin.Default()
	useCase := new(usecase.AuthUseCaseMock)

	router.POST("/api/endpoint", NewAutoMiddleware(useCase), func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	// No Auth Header request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/endpoint", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Empty Auth header request
	w = httptest.NewRecorder()
	req.Header.Set("Authorization", "")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Bearer Auth Header without token request
	w = httptest.NewRecorder()
	useCase.On("ParseToken", "").Return(&models.User{}, auth.ErrInvalidAccessToken)
	req.Header.Set("Authorization", "Bearer ")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Valid Auth Header
	w = httptest.NewRecorder()
	useCase.On("ParseToken", "token").Return(&models.User{}, nil)
	req.Header.Set("Authorization", "Bearer token")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
