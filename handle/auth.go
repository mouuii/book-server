package handle

import (
	"github.com/labstack/echo"
	"net/http"
)

type RegisterForm struct {
	Username string `json:"username" validate:"required,min=5,max=64"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func UserLogin(c echo.Context) error {
	return c.String(http.StatusOK, "login")
}

func UserRegister(c echo.Context) error {

	registerForm := new(RegisterForm)
	if err := c.Bind(registerForm);err != nil {
		return err
	}
	if err := c.Validate(registerForm);err != nil {
		return err
	}
	return c.JSON(http.StatusOK, registerForm)
}