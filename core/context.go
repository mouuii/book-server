package core

import (
	"github.com/codegangsta/inject"
	"net/http"
	"reflect"
)

type Context interface {
	inject.Injector
	Next();
}

type context struct {
	inject.Injector
	Handle  Handle
	Request *http.Request
	Writer  http.ResponseWriter
}

func (this *context) Next() {
	this.run()
}

func (this *context) run() {
	vals, err := this.Invoke(this.Handle)
	if err != nil {
		panic(err)
	}

	ev := this.Get(reflect.TypeOf(ReturnHandler(nil)))
	returnHandle := ev.Interface().(ReturnHandler)
	if len(vals) > 0 {
		returnHandle(this.Writer, this.Request, vals)
	}
}
