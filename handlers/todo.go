package handlers

import (
	"my-todo-app/mock"
	"my-todo-app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var todoRepository = mock.NewMockTodoRepository()

func CreateTodoList(c *gin.Context) {

	var todoList models.TodoList
	if err := c.ShouldBindJSON(&todoList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todoList.CreatedAt = time.Now()
	todoList.UpdatedAt = time.Now()

	if err := todoRepository.CreateTodoList(todoList); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create todo list"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo list created successfully"})
}
