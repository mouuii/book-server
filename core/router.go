package core

import (
	"github.com/codegangsta/inject"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
)

type Router struct {
	inject.Injector
	router *mux.Router
}

type Handle interface{}

func (this *Router) Handle(method string, path string, handle Handle) {
	this.validateHandle(handle)
	this.router.HandleFunc(path, this.wrapHandle(handle)).Methods(method)
}

func (this *Router) GET(path string, handle Handle) {
	this.Handle("GET", path, handle)
}

func (this *Router) Init() {
	if this.router != nil {
		return
	}
	router := mux.NewRouter()
	this.router = router
}

func (this *Router) Register(handle Handle) {
	this.validateHandle(handle)
	_, err := this.Invoke(handle)
	if err != nil {
		panic(err)
	}
}

func (this *Router) Group(prefix string, handle func(*Router)) {
	muxRoute := this.router.PathPrefix(prefix).Subrouter()
	router := &Router{Injector: inject.New(), router: muxRoute}
	router.SetParent(this)
	handle(router)
}

func (this *Router) createContext(writer http.ResponseWriter, request *http.Request, handle Handle) *context {
	c := &context{
		Injector: inject.New(),
		Writer:   writer,
		Request:  request,
		Handle:   handle,
	}
	c.SetParent(this)
	c.MapTo(c, (*Context)(nil))
	c.MapTo(writer, (*http.ResponseWriter)(nil))
	c.Map(request)
	return c
}

func (this *Router) validateHandle(handle Handle) {
	if reflect.TypeOf(handle).Kind() != reflect.Func {
		panic("handler must be a callable func")
	}
}

func (this *Router) wrapHandle(handle Handle) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		c := this.createContext(writer, request, handle)
		c.Next();
	}
}

func (this *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	this.router.ServeHTTP(w, req)
}
