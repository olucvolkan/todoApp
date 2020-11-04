package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/olucvolkan/todoApp/api/models"
	"github.com/olucvolkan/todoApp/api/service"
	"net/http"
)

// GetTodos godoc
// @Summary List all todos
// @Description List todo list
// @Tags todos
// @Accept  json
// @Produce  json
// @Router /todo-list [get]
func GetTodos(c *gin.Context) {
	var todo []models.Todo
	err := models.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, service.CreateTodoListSchema(todo))
	}
}

// GetTodos godoc
// @Summary Create a todo
// @Description  Create a todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param tasks body models.Todo  true "todo info"
// @Router /create-todo [post]
func CreateATodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.CreateATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, todo)
	}
}

// DeleteTodo godoc
// @Summary Delete  a todo
// @Description  Delete  a todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param tasks body models.DeleteTodoRequest  true "true"
// @Router /delete-todo [post]
func DeleteATodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.DeleteATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.AbortWithStatus(http.StatusNoContent)
	}
}

// UpdateTodo godoc
// @Summary Update   a todo
// @Description  Update  a todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param tasks body models.UpdateTodoRequest true "todo info"
// @Router /update-todo [post]
func UpdateATodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.UpdateATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
