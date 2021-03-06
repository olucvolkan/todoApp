package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/olucvolkan/todoApp/api/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(Cors())
	router.GET("/todo-list/", controllers.GetTodos)
	router.POST("/create-todo/", controllers.CreateATodo)
	router.POST("/delete-todo/", controllers.DeleteATodo)
	router.POST("/update-todo/", controllers.UpdateATodo)

	return router
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
		c.Next()
	}
}
