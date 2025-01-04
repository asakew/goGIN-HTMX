package routes

import (
	"github.com/gin-gonic/gin"
	"goGIN-HTMX/internal/handlers"
	"net/http"
)

func HTMLRendering(e *gin.Engine) {
	e.LoadHTMLGlob("web/templates/*")
	e.Static("/files", "web/files")

	e.GET("/", func(c *gin.Context) {
		todos := handlers.ReadToDoList()
		c.HTML(http.StatusOK, "index.html", gin.H{"todos": todos})
	})
}
