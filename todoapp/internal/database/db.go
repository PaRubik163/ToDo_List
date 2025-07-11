package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitSqlite(){
	DB, _ = sql.Open("sqlite3", "./ToDo.db")
}
