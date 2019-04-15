# book-server

> A Simple Framework for Dependency Injection


### example

```go
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
	
	// return value (int, string),int will be return by http status
	router.GET("/404", func() (int, string) {
		return 404, "not fond"
	})

	router.GET("/error", func() error {
		return errors.New("good")
	})

	// auto inject http Request
	router.GET("/inject", func(r *http.Request) string {
		return r.Host
	}) 

	// auto inject App container
	router.GET("/app", func(app *core.App) string {
		return fmt.Sprintf("%v",app.Env)
	})

	app.Run(":8082")

}
```

### config

config file in root project named `conf.yml`

```yaml
env: development
database:
  dialect: postgres
  database_name: postgres
  username: postgres
  password: 123456
  host: 127.0.0.1
  port: 5432
```


### Dependency

- github.com/codegangsta/inject
- github.com/jinzhu/gorm
- github.com/spf13/viper
- github.com/gorilla/mux




### Happy Use ~

