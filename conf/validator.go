package conf

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gopkg.in/go-playground/validator.v9"
)

type AppValidator struct {
	validator *validator.Validate
}

func (v *AppValidator) Validate(i interface{}) error {
	err := v.validator.Struct(i)
	if err == nil {
		return nil
	}
	validateErrs := err.(validator.ValidationErrors)


	log.Printf("%v",validateErrs)
	return validateErrs
}

func addValidator(app *echo.Echo) {
	app.Validator = &AppValidator{validator: validator.New()}
}
