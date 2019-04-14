package core

import (
	"flag"
	"github.com/codegangsta/inject"
	"log"
	"net/http"
	"reflect"
)

var (
	Env = flag.String("env", "dev", "server run mod")
)

type App struct {
	inject.Injector
	Router *Router
	DB     *Database
	Env    string
	Conf   *Config
}

func Initialize() App {

	config := &Config{File: "./conf.yml",}
	if err := config.Init(); err != nil {
		panic(err)
	}

	db := &Database{}
	db.Init()

	router := &Router{}
	router.Init()

	return App{
		DB:     db,
		Router: router,
		Env:    *Env,
		Conf:   config,
	}
}

func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}


type Handle interface{}

func ValidateHandle(handle Handle) {
	if reflect.TypeOf(handle).Kind() != reflect.Func {
		panic("handler must be a callable func")
	}
}

func WrapHandle(handle Handle) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
