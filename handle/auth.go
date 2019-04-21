package handle

import (
	"github.com/labstack/echo"
	"github.com/wowiwj/book-server/handle/context"
	"github.com/wowiwj/book-server/handle/service"
	"github.com/wowiwj/book-server/model"
	"net/http"
)

type RegisterForm struct {
	Username string `json:"username" validate:"required,min=5,max=64"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
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
		return ctx.Failed(http.StatusBadRequest, "用户已存在")
	}

	user.Email = registerForm.Email
	user.Password = registerForm.Password
	user.Name = registerForm.Username

	if _, err := user.Create(); err != nil {
		return ctx.String(http.StatusBadRequest, "注册失败")
	}

	token, err := service.CreateToken(user.ID)
	if err != nil {
		return err
	}
	return ctx.Success(http.StatusOK, AuthResponse{
		Username: user.Name,
		Email:    user.Email,
		Token:    token,
	})
}
