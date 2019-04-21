package app

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type AppContext struct {
	echo.Context
	DB *gorm.DB
}

func (ctx *AppContext) Success(status int, i interface{}) error {
	return ctx.JSON(status, map[string]interface{}{
		"code": status,
		"data": i,
	})
}

func (ctx *AppContext) Failed(status int, i interface{}) error {

	response := ErrorResponse{
		Code:    status,
		Message: i,
	}
	return ctx.JSON(status, response)
}

func (ctx AppContext) Validate(i interface{}) error {
	if err := ctx.Bind(i); err != nil {
		return err
	}
	if err := ctx.Context.Validate(i); err != nil {
		return err
	}
	return nil
}
