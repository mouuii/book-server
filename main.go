package main

import (
	"github.com/wowiwj/book-server/core"
	"net/http"
)


func main() {

	app := core.Initialize()
	router := app.Router

	router.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("index"))
	})

	app.Run(":8082")

}
