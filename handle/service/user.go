package service

import (
	"github.com/wowiwj/book-server/app"
	"github.com/wowiwj/book-server/handle/form"
	"github.com/wowiwj/book-server/model"
)

func CreateUser(form form.RegisterForm) (u *model.User, err error) {
	db := app.GetDB()
	var user model.User

	if ! db.First(&user, "email = ?", form.Email).RecordNotFound() {
		return nil, app.NewApiError("用户已存在")
	}

	user.Email = form.Email
	user.Password = form.Password
	user.Name = form.Username

	if _, err := user.Create(); err != nil {
		return nil, app.NewApiError("注册失败")
	}
	return &user, nil
}
