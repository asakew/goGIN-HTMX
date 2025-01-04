package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HTMLRendering(e *gin.Engine) {
	e.LoadHTMLGlob("web/templates/*")
	e.Static("/files", "web/files")

	// Routes
	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name": "Awesome",
		})
	})
}
