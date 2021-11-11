package config

import (
	"github.com/panjf2000/ants/v2"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"os"
)

func InitRuntime() {
	// init logger
	initLogger(Runtime.Log)

	// init pool
	initPool()

	// init Cron
	initCron()

}

// initCron init cron instance
func initCron() {
	Runtime.Cron = cron.New(cron.WithParser(cron.NewParser(
		cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)))
}

// initPool init pool instance
func initPool() {
	if Runtime.Application.Pool.Size == 0 {
		Runtime.Pool, _ = ants.NewPool(10000)
	} else {
		Runtime.Pool, _ = ants.NewPool(Runtime.Application.Pool.Size)
	}
}

// initLogger init logger
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
