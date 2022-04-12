package util

import (
	"clean-arch/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseDriver string

const (
	MySQL DatabaseDriver = "mysql"
)

type DatabaseConnection struct {
	Driver DatabaseDriver

	MySQL *gorm.DB
}

func NewConnectionDatabase(config *config.AppConfig) *DatabaseConnection {
	var db DatabaseConnection

	switch config.Database.Driver {
	case "MySQL":
		db.Driver = MySQL
		db.MySQL = newMySQL(config)
	}

	return &db

}

func newMySQL(config *config.AppConfig) *gorm.DB {
	db_url := config.Database.Db_url
	db, err := gorm.Open(mysql.Open(db_url), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func (db *DatabaseConnection) CloseConnection() {
	if db.MySQL != nil {
		db, _ := db.MySQL.DB()
		db.Close()
	}
}
