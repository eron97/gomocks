package handlers

import (
	"net/http"

	"github.com/eron97/gomocks.git/services"
	"github.com/gin-gonic/gin"
)

type TodoListHandler struct {
	TodoListService services.TodoListService
}

func (h *TodoListHandler) GetTodoList(c *gin.Context) {
	todoList, err := h.TodoListService.GetTodoList(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todoList)
}
