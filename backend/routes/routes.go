package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/olucvolkan/todoApp/controllers"
)
func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("todo-list", controllers.GetTodos)
		v1.POST("create-todo", controllers.CreateATodo)
		v1.POST("delete-todo", controllers.DeleteATodo)
		v1.PUT("update-todo", controllers.UpdateATodo)
		
	}

	return r
}
