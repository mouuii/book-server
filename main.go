package main

import (
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github.com/wowiwj/book-server/conf"
)

func main() {

	app := echo.New()
	if err := conf.Init(app); err != nil {
		app.Logger.Fatal(err)
		return
	}
	port := viper.GetString("port")
	app.Logger.Fatal(app.Start(":" + port))
}
