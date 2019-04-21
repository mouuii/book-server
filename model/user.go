package model

import "github.com/wowiwj/book-server/app"

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
