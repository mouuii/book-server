package conf

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	db, err := gorm.Open(config.Dialect, databaseUrl)
	if err != nil {
		return err
	}
	this.DB = db
	return nil
}
