package main

import (
	"errors"
	"fmt"
	"github.com/wowiwj/book-server/core"
	"net/http"
)

func main() {

	app := core.Initialize()
	router := app.Router

	router.GET("/", func() string {
		return "hello"
	})

	router.GET("/404", func() (int, string) {
		return 404, "not fond"
	})

	router.GET("/error", func() error {
		return errors.New("good")
	})

	router.GET("/inject", func(r *http.Request) string {
		fmt.Println(r.Host)
		return r.Host
	})

	router.GET("/app", func(app *core.App) string {
		return fmt.Sprintf("%v", app.Env)
	})

	app.Run(":8082")

}
