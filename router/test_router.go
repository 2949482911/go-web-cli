package router

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func init() {
	AddRouter(initTestRouter)
}

func initTestRouter(party iris.Party) {
	party.Get("/test", func(context *context.Context) {
		context.WriteString("hello test")
	})
}
