package conf

import (
	"github.com/jinzhu/gorm"
	"github.com/wowiwj/book-server/model"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(model.User{})
}
