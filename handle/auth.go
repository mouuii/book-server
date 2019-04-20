package handle

import (
	"github.com/labstack/echo"
	"github.com/wowiwj/book-server/handle/context"
	"github.com/wowiwj/book-server/model"
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
	if err := c.Bind(registerForm); err != nil {
		return err
	}
	if err := c.Validate(registerForm); err != nil {
		return err
	}

	ctx := c.(context.AppContext)
	db := ctx.DB

	var user model.User

	if ! db.First(&user, "email = ?", registerForm.Email).RecordNotFound() {
		return ctx.String(http.StatusOK, "用户已存在")
	}

	user.Email = registerForm.Email
	user.Password = registerForm.Password
	user.Name = registerForm.Username

	if err := db.Create(&user).Error; err != nil {
		return ctx.String(http.StatusOK, "注册失败")
	}

	return c.JSON(http.StatusOK, registerForm)
}
