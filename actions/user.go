package actions

import "github.com/wowiwj/book-server/core"

type UserAction struct {
	*core.Action
}

func (this *UserAction) Index() string {
	return "user index"
}
