package router

import (
	"github.com/labstack/echo"
	"github.com/wowiwj/book-server/app"
	"github.com/wowiwj/book-server/handle"
)

type HandleFunc func(ctx app.AppContext) error

func w(handleFunc HandleFunc) echo.HandlerFunc {
	return func(i echo.Context) error {
		ctx := i.(app.AppContext)
		return handleFunc(ctx)
	}
}

func Init(router *echo.Echo)  {

	auth := router.Group("/auth")
	{
		auth.POST("/login", w(handle.UserLogin))
		auth.POST("/register",w(handle.UserRegister))
	}
}