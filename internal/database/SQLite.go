package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var SQLiteDB *sql.DB

func DBSQLite() {
	var err error

	SQLiteDB, err = sql.Open("sqlite3", "./internal/database/SQLite.db")
	if err != nil {
		log.Fatal(err)
	}

	stmt := `
    CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    );`
	_, err = SQLiteDB.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Database initialized")
}

func CloseDatabase() {
	err := SQLiteDB.Close()
	if err != nil {
		return
	}
}
