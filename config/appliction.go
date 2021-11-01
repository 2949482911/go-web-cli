package config

import (
	"github.com/kataras/iris/v12"
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
