package conf

import "github.com/wowiwj/book-server/model"

func migrate(db *Database) {
	db.DB.AutoMigrate(model.User{})
}
