package main

import (
	"goGIN-HTMX/internal/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	// Database
	database.DBSQLite()

	// HTML Rendering
	e.LoadHTMLGlob("web/templates/*")
	e.Static("/files", "web/files")

	// Routes
	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name": "Awesome",
		})
	})

	// Start Server
	log.Printf("Starting server on port http://localhost:8080")
	err := e.Run(":8080")
	if err != nil {
		return
	}
}
