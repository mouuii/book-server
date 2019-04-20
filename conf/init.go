package conf

import (
	"github.com/labstack/echo"
	context2 "github.com/wowiwj/book-server/handle/context"
	"github.com/wowiwj/book-server/router"
)

type Ctx struct {
	conf *Config
	DB   *Database
	App  *echo.Echo
}



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

	// register validate
	addValidator(app)
	registerError(app)
	
	app.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			cc := context2.AppContext{context,db.DB}
			return handlerFunc(cc)
		}
	})
	migrate(db)
	return nil
}
