package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/olucvolkan/todoApp/models"
	"github.com/olucvolkan/todoApp/service"
	"net/http"
)

func GetTodos(c *gin.Context) {
	var todo []models.Todo
	err := models.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, 	service.CreateTodoListSchema(todo))
	}
}

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
