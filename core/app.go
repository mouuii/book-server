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

func Initialize() *App {

	flag.Parse();
	app := &App{
		Injector: inject.New(),
		Env: *Env,
	}

	app.initRouter()
	app.initConfig()
	app.initDatabase()

	// inject app and default handle
	app.Map(app)
	app.Map(defaultReturnHandle())
	return app
}

func (app *App) initRouter() *Router {
	router := &Router{Injector: inject.New()}
	router.Init()
	router.SetParent(app)
	app.Map(router)
	app.Router = router
	return router
}

func (app *App) initConfig() *Config {
	config := &Config{Injector: inject.New(), File: "./conf.yml",}
	if err := config.Init(); err != nil {
		panic(err)
	}
	config.SetParent(app)
	app.Map(config)
	app.Conf = config
	return config
}

func (app *App)initDatabase() *Database  {
	db := &Database{Injector: inject.New()}
	db.Init()
	db.SetParent(app)
	app.Map(db)
	app.DB = db
	return db
}

func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app))
}

func (app *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	app.Router.ServeHTTP(w, req)
}
