package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/olucvolkan/todoApp/models"
	"net/http"
)

func GetTodos(c *gin.Context) {
	var todo []models.Todo
	err := models.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
})
