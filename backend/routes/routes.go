package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/olucvolkan/todoApp/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/todo-list/", controllers.GetTodos)
	router.POST("/create-todo/", controllers.CreateATodo)
	router.POST("/delete-todo/", controllers.DeleteATodo)
	router.PUT("/update-todo/", controllers.UpdateATodo)

	return router
}
