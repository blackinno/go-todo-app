package routers

import (
	"backend.api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	todoGroup := r.Group("/")
	{
		todoGroup.GET("/todos", controllers.ListTodo)
		todoGroup.POST("/todo/create", controllers.CreateTodo)
		todoGroup.PUT("/todo/:id", controllers.UpdateTodo)
		todoGroup.DELETE("/todo/:id", controllers.DeleteTodo)
	}
}
