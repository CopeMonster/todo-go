package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-go/auth"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type signInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) SignUp(ctx *gin.Context) {
	inp := new(signInput)

	if err := ctx.BindJSON(inp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.useCase.SignUp(ctx.Request.Context(), inp.Username, inp.Password); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

type signInResponse struct {
	Token string `json:"token"`
}

func (h *Handler) SignIn(ctx *gin.Context) {
	inp := new(signInput)

	if err := ctx.BindJSON(inp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.useCase.SignIn(ctx.Request.Context(), inp.Username, inp.Password)

	if err != nil {
		if errors.Is(err, auth.ErrUserNotFound) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, signInResponse{
		Token: token,
	})
}
