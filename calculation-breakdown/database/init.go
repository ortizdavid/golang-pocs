package database

import (
	"database/sql"
    _ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
)

func InitDB(filepath string) *sql.DB {
    db, err := sql.Open("sqlite3", filepath)
    if err != nil {
        log.Fatal("Error creating database:", err)
    }

    db.SetMaxOpenConns(1)     
    db.SetConnMaxLifetime(time.Minute * 5)

    if err = db.Ping(); err != nil {
        log.Fatal("Error connecting to SQLite:", err)
    }

    createTables(db)
    return db
}

func createTables(db *sql.DB) {
	// read from tables.sql
    query, err := os.ReadFile("./database/tables.sql")
	if err != nil {
		log.Fatal("Error reading tables.sql:", err)
	}
    _, err = db.Exec(string(query))
    if err != nil {
        log.Fatal("error while creating tables:", err)
    }
}