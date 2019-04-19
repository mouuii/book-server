package handle

import (
	"github.com/labstack/echo"
	"net/http"
)

func UserLogin(c echo.Context) error {
	return c.String(http.StatusOK, "login")
}

func UserRegister(c echo.Context) error {
	return c.String(http.StatusOK, "register")
}