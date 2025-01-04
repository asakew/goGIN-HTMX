package main

import (
	"github.com/gin-gonic/gin"
	"goGIN-HTMX/internal/database"
	"goGIN-HTMX/internal/routes"
	"log"
)

func main() {
	e := gin.Default()

	// Database
	database.DBSQLite()
	defer database.CloseDatabase()

	// Routes
	routes.HTMLRendering(e)
	routes.TodoRoutes(e)

	// Start Server
	log.Printf("Starting server on port http://localhost:8080")
	err := e.Run(":8080")
	if err != nil {
		return
	}
}
