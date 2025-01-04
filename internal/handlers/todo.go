package handlers

import (
	"database/sql"
	"goGIN-HTMX/internal/database"
	"goGIN-HTMX/internal/models"
)

func CreateToDo(title string, status string) (int64, error) {
	result, err := database.SQLiteDB.Exec("INSERT INTO todos (title, status) VALUES (?, ?)", title, status)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func DeleteToDo(id int64) error {
	_, err := database.SQLiteDB.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}

func ReadToDoList() []models.ToDo {
	rows, err := database.SQLiteDB.Query("SELECT id, title, status FROM todos")
	if err != nil {
	}
	defer func(rows *sql.Rows) {
		rows.Close()
	}(rows)

	todos := make([]models.ToDo, 0)
	for rows.Next() {
		var todo models.ToDo
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Status)
		if err != nil {
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		// Handle error
	}
	return todos
}

//func UpdateToDo(id int64, title string, status string) error {
//	_, err := database.SQLiteDB.Exec("UPDATE todos SET title = ?, status = ? WHERE id = ?", title, status, id)
//	return err
//}
