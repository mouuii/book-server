package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	DB *gorm.DB
}

var database *Database

func InitDB(dialect, databaseUrl string) error {
	if database != nil {
		return nil
	}
	db, err := gorm.Open(dialect, databaseUrl)
	if err != nil {
		return err
	}
	database = &Database{
		DB: db,
	}
	return nil
}

func GetDatabase() *Database {
	return database
}

func GetDB() *gorm.DB {
	return database.DB
}
