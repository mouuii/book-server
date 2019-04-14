package core

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	router *mux.Router
}

func (this *Router)GET(path string, f func(http.ResponseWriter,
	*http.Request))  {
	this.router.Methods("GET").HandlerFunc(f)
}

func (this *Router) Init() {
	if this.router != nil {
		return
	}
	router := mux.NewRouter()
	this.router = router
}

func (this *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	this.router.ServeHTTP(w,req)
}
