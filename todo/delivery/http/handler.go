package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"todo-go/auth"
	"todo-go/models"
	"todo-go/todo"
)

type Todo struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Done         bool      `json:"done"`
	CreatedTime  time.Time `json:"createdTime"`
	ModifiedTime time.Time `json:"modifiedTime"`
}

type Handler struct {
	useCase todo.UseCase
}

func NewHandler(useCase todo.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) GetTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	user := ctx.MustGet(auth.CtxUserKey).(*models.User)

	td, err := h.useCase.GetTodo(ctx.Request.Context(), user, id)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, td)
}

type getTodosResponse struct {
	Todos []*Todo `json:"todos"`
}

func (h *Handler) GetTodos(ctx *gin.Context) {
	user := ctx.MustGet(auth.CtxUserKey).(*models.User)

	todos, err := h.useCase.GetTodos(ctx.Request.Context(), user)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, &getTodosResponse{
		Todos: toTodos(todos),
	})
}

type createTodoInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *Handler) CreateTodo(ctx *gin.Context) {
	inp := new(createTodoInput)

	if err := ctx.BindJSON(inp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := ctx.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.CreateTodo(ctx.Request.Context(), user, inp.Title, inp.Description); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

type updateTodoInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func (h *Handler) UpdateTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	inp := new(updateTodoInput)

	if err := ctx.BindJSON(inp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := ctx.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.UpdateTodo(ctx.Request.Context(), user, id, inp.Title, inp.Description, inp.Done); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	user := ctx.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.useCase.DeleteTodo(ctx.Request.Context(), user, id); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func toTodos(tds []*models.Todo) []*Todo {
	out := make([]*Todo, len(tds))

	for i, t := range tds {
		out[i] = toTodo(t)
	}

	return out
}

func toTodo(td *models.Todo) *Todo {
	return &Todo{
		ID:           td.ID,
		Title:        td.Title,
		Description:  td.Description,
		Done:         td.Done,
		CreatedTime:  td.CreatedTime,
		ModifiedTime: td.ModifiedTime,
	}
}
