package core

import (
	"fmt"
	"github.com/codegangsta/inject"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	inject.Injector
	router *mux.Router
}

func (this *Router) Handle(method string,path string,handle Handle) {
	ValidateHandle(handle)


	this.router.HandleFunc(path,func(writer http.ResponseWriter, request *http.Request) {
		c := this.createContext(writer,request)
		vals,err := c.Invoke(handle)
		if err != nil {
			panic(err)
		}

		fmt.Println(vals)

		returnHandle := defaultReturnHandle()
		if len(vals) > 0 {
			returnHandle(writer,request,vals)
		}

	}).Methods(method)
}


func (this *Router) GET1(path string, f Handle) {
	this.Handle("GET",path,f)
}

func (this *Router) GET(path string, f func(http.ResponseWriter,
	*http.Request)) *Router {
	this.router.Methods("GET").HandlerFunc(f)
	return this
}

func (this *Router) Init() {
	if this.router != nil {
		return
	}
	router := mux.NewRouter()
	this.router = router
}

func (this *Router) createContext(writer http.ResponseWriter, request *http.Request) *context {
	c := &context{inject.New()}
	c.SetParent(this)
	c.MapTo(c, (*Context)(nil))
	c.MapTo(writer,(*http.ResponseWriter)(nil))
	c.Map(request)
	return c
}

func (this *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	this.router.ServeHTTP(w, req)
}
