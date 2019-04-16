package routes

import (
	"errors"
	"fmt"
	. "github.com/wowiwj/book-server/actions"
	"github.com/wowiwj/book-server/core"
	"net/http"
)


func Api(router *core.Router) {

	router.Group("/users", func(router *core.Router) {
		router.GET("",(&UserAction{}).Index)
		router.GET("/:id","UserAction@show")
	})


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
}