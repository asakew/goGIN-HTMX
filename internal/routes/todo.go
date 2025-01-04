package routes

import (
	"github.com/gin-gonic/gin"
	"goGIN-HTMX/internal/handlers"
	"net/http"
	"strconv"
)

func TodoRoutes(e *gin.Engine) {
	e.POST("/todos", func(c *gin.Context) {
		title := c.PostForm("title")
		status := c.PostForm("status")
		id, _ := handlers.CreateToDo(title, status)

		c.HTML(http.StatusOK, "task.html", gin.H{
			"title":  title,
			"status": status,
			"id":     id,
		})
	})

	e.DELETE("/todos/:id", func(c *gin.Context) {
		param := c.Param("id")
		id, _ := strconv.ParseInt(param, 10, 64)

		err := handlers.DeleteToDo(id)
		if err != nil {
			return
		}
	})
}
