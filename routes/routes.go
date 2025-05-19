package routes

import (
	"go-todo/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {

	// Check if /todos is defined here
	r.POST("/todo", controllers.CreateTodo)
	r.GET("/todos", controllers.GetTodos)
}
