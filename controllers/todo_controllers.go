package controllers

import (
	"go-todo/config"
	"go-todo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON((http.StatusBadRequest), gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&todo)
	c.JSON(http.StatusCreated, gin.H{"data": todo})
}

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	config.DB.Find(&todos)

	c.JSON(http.StatusOK, gin.H{"data": todos})
}
