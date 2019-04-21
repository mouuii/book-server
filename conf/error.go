package conf

import (
	"github.com/labstack/echo"
	"github.com/wowiwj/book-server/app"
)

func registerError(app *echo.Echo) {
	app.HTTPErrorHandler = httpErrorHandler
}

func httpErrorHandler(err error, c echo.Context) {
	parser := app.NewErrorParser(err, c)
	parser.Process().Response()

}
