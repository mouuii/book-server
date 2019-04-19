package conf

import "github.com/wowiwj/book-server/model"

func Migrate(db *Database) {
	db.DB.AutoMigrate(model.User{})
}
