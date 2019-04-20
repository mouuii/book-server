package conf

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
	"strings"
)

type AppValidator struct {
	validator *validator.Validate
}

func (v *AppValidator) Validate(i interface{}) error {
	err := v.validator.Struct(i)
	if err == nil {
		return nil
	}
	return err
}

func (v *AppValidator) TransError(err validator.ValidationErrors) validator.ValidationErrorsTranslations {
	zh_cn := zh.New()
	uni := ut.New(zh_cn, zh_cn)
	trans, _ := uni.GetTranslator("en")

	_ = zh_translations.RegisterDefaultTranslations(v.validator, trans)
	painTrans := err.Translate(trans)
	return v.convertField(painTrans)
}

func (v *AppValidator) convertField(translate validator.ValidationErrorsTranslations) validator.ValidationErrorsTranslations {

	trans := map[string]string{}

	for k, v := range translate {
		from := strings.Split(k, ".")
		newKey := from[len(from)-1]
		trans[newKey] = v
	}
	return trans
}

func addValidator(app *echo.Echo) {
	v := validator.New()
	app.Validator = &AppValidator{validator: v}
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

}
