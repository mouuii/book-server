package core

import (
	"fmt"
	"github.com/codegangsta/inject"
	"github.com/jinzhu/gorm"
	"log"

	//_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// import _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	inject.Injector
	DB *gorm.DB
}

func (this *Database) Init() {

	config := GetDbConfig()
	databaseUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Host,
		config.Port,
		config.UserName,
		config.Name,
		config.Password,
	)
	log.Println(databaseUrl)
	db, err := gorm.Open(config.Dialect, databaseUrl)
	if err != nil {
		log.Fatalln(err)
	}
	this.DB = db
}
