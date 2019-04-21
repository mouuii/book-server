package app

import (
	"fmt"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type ErrorParser struct {
	e        *echo.Echo
	context  echo.Context
	response *ErrorResponse
	err      error
}

func NewErrorParser(err error, c echo.Context) ErrorParser {
	parser := ErrorParser{
		e:       c.Echo(),
		context: c,
		err:     err,
	}
	return parser
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

func (p *ErrorParser) ParseCustomer() *ErrorResponse {
	response := &ErrorResponse{}
	if err, ok := p.err.(validator.ValidationErrors); ok {
		response.Code = 422
		v := p.e.Validator.(*AppValidator)
		response.Message = v.TransError(err)
		return response
	}

	if err,ok := p.err.(*ApiError);ok {
		response.Code = err.status
		response.Message = err.message
		return response
	}

	return nil
}

func (p *ErrorParser) parseDefault() *ErrorResponse {
	response := &ErrorResponse{}
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

type ApiError struct {
	status  int
	message interface{}
}

func (err ApiError) Error() string {
	return fmt.Sprintf("status: %d, error: %v", err.Status, err.message)
}

func NewApiError(message interface{}) *ApiError {
	return &ApiError{
		status:  400,
		message: message,
	}
}

func (err *ApiError) Status(status int) *ApiError {
	err.status = status
	return err
}

