package config

import (
	"github.com/sirupsen/logrus"
	"os"
)

func InitRuntime() {
	// init logger
	initLogger(Runtime.Log)
}

func initLogger(logger *logrus.Logger) {
	Runtime.Log = logger
	switch Runtime.Application.Logger.Level {
	case "info":
		Runtime.Log.SetLevel(logrus.InfoLevel)
		break
	case "warn":
		Runtime.Log.SetLevel(logrus.WarnLevel)
		break
	case "debug":
		Runtime.Log.SetLevel(logrus.DebugLevel)
		break
	case "error":
		Runtime.Log.SetLevel(logrus.ErrorLevel)
		break
	case "panic":
		Runtime.Log.SetLevel(logrus.PanicLevel)
		break
	case "fatal":
		Runtime.Log.SetLevel(logrus.FatalLevel)
		break
	default:
		Runtime.Log.SetLevel(logrus.DebugLevel)
		break
	}

	if Runtime.Application.Logger.Format == "json" {
		Runtime.Log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		// The TextFormatter is default, you don't actually have to do this.
		Runtime.Log.SetFormatter(&logrus.TextFormatter{})
	}
	Runtime.Log.Out = os.Stdout
}
