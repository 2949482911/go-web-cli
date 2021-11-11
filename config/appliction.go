package config

import (
	"github.com/kataras/iris/v12"
	"github.com/panjf2000/ants/v2"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

// Runtime runtime object
var Runtime *App

func init() {
	if Runtime == nil {
		Runtime = GetEngine()
	}
}

// App App 容器
type App struct {
	Engine      *iris.Application  // application
	Router      []func(iris.Party) // router list
	Application *Application       // application
	Log         *logrus.Logger     // logger
	Pool        *ants.Pool         // pool
	Cron        *cron.Cron         // Cron
}

// GetEngine 获取引擎一个默认引擎
func GetEngine() *App {
	return &App{
		Engine: iris.New(),
		Router: make([]func(iris.Party), 0),
		Log:    logrus.New(),
	}
}

// StartEngine start the web engine
func StartEngine() {
	if Runtime.Application.GetHost() == "" {
		Runtime.Application = getDefaultConfig()
	}
	_ = Runtime.Engine.Run(iris.Addr(Runtime.Application.GetHost()))
	return
}

func ExternalFunc(fun ...func()) {
	for _, f := range fun {
		go f()
	}
}
