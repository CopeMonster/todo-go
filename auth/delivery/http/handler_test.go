package http

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-go/auth/usecase"
)

func TestSignUp(t *testing.T) {
	router := gin.Default()
	useCase := new(usecase.AuthUseCaseMock)

	RegisterHTTPEndpoints(router, useCase)

	signUpBody := &signInput{
		Username: "testuser",
		Password: "testpass",
	}

	body, err := json.Marshal(signUpBody)
	assert.NoError(t, err)

	useCase.On("SignUp", signUpBody.Username, signUpBody.Password).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/sign-up", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSignIn(t *testing.T) {
	router := gin.Default()
	useCase := new(usecase.AuthUseCaseMock)

	RegisterHTTPEndpoints(router, useCase)

	signUpBody := &signInput{
		Username: "testuser",
		Password: "testpass",
	}

	body, err := json.Marshal(signUpBody)
	assert.NoError(t, err)

	useCase.On("SignIn", signUpBody.Username, signUpBody.Password).Return("jwt", nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/sign-in", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"token\":\"jwt\"}", w.Body.String())
}
