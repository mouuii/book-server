package router

import (
	"github.com/labstack/echo"
	"github.com/wowiwj/book-server/handle"
)

func Init(router *echo.Echo)  {

	auth := router.Group("/auth")
	{
		auth.GET("/login", handle.UserLogin)
		auth.POST("/register",handle.UserRegister)
	}


}