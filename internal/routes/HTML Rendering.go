package routes

import (
	"github.com/gin-gonic/gin"
)

func HTMLRendering(e *gin.Engine) {
	e.LoadHTMLGlob("web/templates/*")
	e.Static("/files", "web/files")
}
