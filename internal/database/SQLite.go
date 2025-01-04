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

	_, err = SQLiteDB.Exec(`
    CREATE TABLE IF NOT EXISTS todos (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT,
            status TEXT
    );`)

	log.Printf("Database initialized")

	if err != nil {
		log.Fatal(err)
	}
}

func CloseDatabase() {
	err := SQLiteDB.Close()
	if err != nil {
		return
	}
}
