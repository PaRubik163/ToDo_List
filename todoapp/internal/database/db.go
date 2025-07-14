package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

var DB *sql.DB

func InitSqlite() error {
	var err error
	DB, err = sql.Open("sqlite3", "./ToDo.db")

	if err != nil{
		return err
	}

	logrus.Warn("Ошибка открытия базы данных")
	return nil
}
