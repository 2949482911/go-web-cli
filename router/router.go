package router

import (
	"github.com/go-web-cli/config"
	"github.com/kataras/iris/v12"
)

// AddRouter 初始化的router
func AddRouter(args ...func(party iris.Party)) {
	for _, f := range args {
		config.Runtime.Router = append(config.Runtime.Router, f)
	}
}

// GetRouter get routers data list
func GetRouter() []func(party iris.Party) {
	return config.Runtime.Router
}

// ExeRouter exe router
func ExeRouter() {
	for _, f := range config.Runtime.Router {
		if config.Runtime.Application.ContextPath != "" {
			f(config.Runtime.Engine.Party(config.Runtime.Application.ContextPath))
		} else {
			f(config.Runtime.Engine.Party("/"))
		}
	}
}
