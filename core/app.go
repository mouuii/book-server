package core

import (
	"flag"
	"github.com/codegangsta/inject"
	"log"
	"net/http"
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
	log.Fatal(http.ListenAndServe(addr, app))
}

func (app *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	app.Router.ServeHTTP(w, req)
}



