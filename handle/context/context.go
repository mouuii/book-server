package context

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type AppContext struct {
	echo.Context
	DB *gorm.DB
}
