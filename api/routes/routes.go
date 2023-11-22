package routes

import (
	"github.com/eron97/gomocks.git/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *handlers.TodoListHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/todolist", handler.GetTodoList)

	return r
}
