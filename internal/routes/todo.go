package routes

import (
	"github.com/gin-gonic/gin"
	"goGIN-HTMX/internal/handlers"
)

func TaskRoutes(e *gin.Engine) {
	e.GET("/", handlers.GetTasks)
	e.POST("/tasks", handlers.AddTask)
	e.DELETE("/tasks/:id", handlers.DeleteTask)
}
