package conf

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/wowiwj/book-server/handle/context"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func registerError(app *echo.Echo) {
	app.HTTPErrorHandler = httpErrorHandler
}

type ErrorParser struct {
	e        *echo.Echo
	context  echo.Context
	response *context.ErrorResponse
	err      error
}

func httpErrorHandler(err error, c echo.Context) {

	parser := ErrorParser{
		e:       c.Echo(),
		context: c,
		err:     err,
	}

	parser.Process().Response()

}

func (p *ErrorParser) Response() {
	c := p.context

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			p.err = c.NoContent(p.response.Code)
		} else {
			p.err = c.JSON(p.response.Code, p.response)
		}
		if p.err != nil {
			p.e.Logger.Error(p.err)
		}
	}
}

func (p *ErrorParser) Process() *ErrorParser {
	response := p.ParseCustomer()
	if response != nil {
		p.response = response
		return p
	}
	p.response = p.parseDefault()
	return p
}

func (p *ErrorParser) ParseCustomer() *context.ErrorResponse {
	response := &context.ErrorResponse{}
	if err, ok := p.err.(validator.ValidationErrors); ok {
		response.Code = 422
		v := p.e.Validator.(*AppValidator)
		response.Message = v.TransError(err)
		return response
	}
	return nil
}

func (p *ErrorParser) parseDefault() *context.ErrorResponse {
	response := &context.ErrorResponse{}
	if he, ok := p.err.(*echo.HTTPError); ok {
		response.Code = he.Code
		response.Message = he.Message
		if he.Internal != nil {
			p.err = fmt.Errorf("%v, %v", p.err, he.Internal)
		}
	} else if p.e.Debug {
		response.Message = p.err.Error()
	} else {
		response.Message = http.StatusText(response.Code)
	}
	return response
}
