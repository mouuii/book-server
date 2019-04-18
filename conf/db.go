package conf

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/gommon/log"
)

type Database struct {
	DB *gorm.DB
}

func (this *Database) Init() error {
	config := GetDbConfig()
	databaseUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Host,
		config.Port,
		config.UserName,
		config.Name,
		config.Password,
	)

	log.Debug(databaseUrl)
	db, err := gorm.Open(config.Dialect, databaseUrl)
	if err != nil {
		return err
	}
	this.DB = db
	return nil
}
