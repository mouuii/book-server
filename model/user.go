package model

import (
	"github.com/wowiwj/book-server/app"
	"github.com/wowiwj/book-server/util"
)

type User struct {
	BaseModel
	Name     string `gorm:"size:64"`
	Email    string `gorm:"size:128;index:idx_user_email"`
	Password string
	Phone    string `gorm:"size:32;index:idx_user_phone"`
	IsAdmin  bool   `gorm:"default:false;not null"`
	Avatar   string
	Bio      string
}

func (u *User) Create() (id uint, err error) {
	db := app.GetDB()
	if err = db.Create(u).Scan(u).Error; err != nil {
		return
	}
	return u.ID, nil
}

func (u *User) IsExist() bool {
	db := app.GetDB()
	return !db.First(&u, "email = ?", u.Email).RecordNotFound()
}

func (u *User) BeforeSave() (err error) {

	secret, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}
	// 加密
	u.Password = secret
	return nil
}

func (u *User) GetUserByEmail(email string) error {
	db := app.GetDB()
	return db.Where("email = ?", email).First(u).Error
}
