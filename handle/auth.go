package handle

import (
	"github.com/labstack/echo"
	"github.com/wowiwj/book-server/app"
	"github.com/wowiwj/book-server/handle/form"
	"github.com/wowiwj/book-server/handle/service"
	"github.com/wowiwj/book-server/model"
	"net/http"
)

type AuthResponse struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func UserLogin(c echo.Context) error {
	return c.String(http.StatusOK, "login")
}

func UserRegister(ctx app.AppContext) error {
	var (
		user  *model.User
		err   error
		token string
	)

	registerForm := new(form.RegisterForm)
	if err = ctx.Validate(registerForm); err != nil {
		return err
	}

	if user, err = service.CreateUser(*registerForm); err != nil {
		return err
	}

	if token, err = service.CreateToken(user.ID); err != nil {
		return err
	}
	return ctx.Success(http.StatusOK, AuthResponse{
		Id:       user.ID,
		Username: user.Name,
		Email:    user.Email,
		Token:    token,
	})
}
