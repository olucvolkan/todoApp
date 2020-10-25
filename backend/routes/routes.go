package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/olucvolkan/todoApp/controllers"
)
func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("todo", controllers.GetTodos)
		
	}

	return r
}
