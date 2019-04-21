package conf

import (
	"github.com/labstack/echo"
	"github.com/wowiwj/book-server/app"
	"github.com/wowiwj/book-server/handle/context"
	"github.com/wowiwj/book-server/router"
)

type Ctx struct {
	conf *Config
	DB   *app.Database
	App  *echo.Echo
}

func Init(e *echo.Echo) error {
	conf := &Config{
		File: "conf.yml",
	}
	if err := conf.InitConfig(); err != nil {
		return err
	}

	if err := InitDB(); err != nil {
		return err
	}

	db := app.GetDB()

	router.Init(e)

	// register validate
	addValidator(e)
	registerError(e)

	e.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			cc := context.AppContext{ctx, db}
			return handlerFunc(cc)
		}
	})
	migrate(db)
	return nil
}
