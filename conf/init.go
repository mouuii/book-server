package conf

import (
	"github.com/labstack/echo"
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
	return nil
}
