package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	e.LoadHTMLGlob("templates/*")

	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name": "Awesome",
		})
	})

	err := e.Run(":8080")
	if err != nil {
		return
	}
}
