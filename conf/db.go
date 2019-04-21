package conf

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/wowiwj/book-server/app"
)

func GetDBUrl(config DbConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Host,
		config.Port,
		config.UserName,
		config.Name,
		config.Password,
	)
}

func InitDB() error {
	config := GetDbConfig()
	return app.InitDB(config.Dialect, GetDBUrl(config))
}
