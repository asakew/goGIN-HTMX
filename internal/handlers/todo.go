package handlers

import (
	"github.com/gin-gonic/gin"
	"goGIN-HTMX/internal/database"
	"goGIN-HTMX/internal/models"
	"net/http"
)

func GetTasks(c *gin.Context) {
	rows, err := database.SQLiteDB.Query("SELECT id, name FROM tasks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to query database"})
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"tasks": tasks,
	})
}

func AddTask(c *gin.Context) {
	name := c.PostForm("name")
	_, err := database.SQLiteDB.Exec("INSERT INTO tasks (name) VALUES (?)", name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to insert into database"})
		return
	}

	// Вернём обновлённый список задач
	GetTasks(c)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	_, err := database.SQLiteDB.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete from database"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Task deleted"})
}
