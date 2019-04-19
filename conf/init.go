package conf

import (
	"github.com/labstack/echo"
	"github.com/wowiwj/book-server/router"
)

func Init(app *echo.Echo) error {
	conf := &Config{
		File: "conf.yml",
	}
	if err := conf.InitConfig(); err != nil {
		return err
	}

	db := &Database{}
	if err := db.Init(); err != nil {
		return err
	}

	router.Init(app)

	return nil
}
