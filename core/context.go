package core

import "github.com/codegangsta/inject"

type Context interface {
	inject.Injector
}

type context struct {
	inject.Injector

}

