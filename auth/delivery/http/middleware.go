package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"todo-go/auth"
)

type AutoMiddleware struct {
	useCase auth.UseCase
}

func NewAutoMiddleware(useCase auth.UseCase) gin.HandlerFunc {
	return (&AutoMiddleware{
		useCase: useCase,
	}).Handle
}

func (m *AutoMiddleware) Handle(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := m.useCase.ParseToken(ctx.Request.Context(), headerParts[1])

	if err != nil {
		status := http.StatusInternalServerError

		if errors.Is(err, auth.ErrInvalidAccessToken) {
			status = http.StatusUnauthorized
		}

		ctx.AbortWithStatus(status)
		return
	}

	ctx.Set(auth.CtxUserKey, user)
}
