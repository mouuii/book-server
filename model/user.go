package model

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
