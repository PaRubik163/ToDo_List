package storage

import (

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"todoapp/internal/models"
)

var DB *gorm.DB

func InitSqlite() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("ToDo.db"), &gorm.Config{})

	if err != nil{
		logrus.Warn("Отствует возможность открыть бд ", err)
		return err
	}
	logrus.Info("База данных инициализирована")

	err = DB.AutoMigrate(&models.Todo{})
	if err != nil{
		return nil
	}

	logrus.Info("Создана таблица")
	return nil
}
