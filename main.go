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

	router.GET1("/", func() string {
		return "hello"
	})

	router.GET1("/404", func() (int, string) {
		return 404, "not fond"
	})

	router.GET1("/error", func() error {
		return errors.New("good")
	})

	router.GET1("/inject", func(r *http.Request) string {
		fmt.Println(r.Host)
		return r.Host
	})

	app.Run(":8082")

}
