package main

import (
	"github.com/wowiwj/book-server/core"
	"github.com/wowiwj/book-server/routes"
)

func main() {

	app := core.Initialize()

	app.Router.Register(routes.Api)

	app.Run(":8082")
}
