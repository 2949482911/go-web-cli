package config

import (
	"bytes"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strconv"
)

// Application web app configuration
type Application struct {
	Host        string      `json:"host" yaml:"host"`
	Port        int         `json:"port" yaml:"port"`
	ContextPath string      `json:"context_path" yaml:"context-path"`
	Logger      Logger      `json:"logger" yaml:"logger"`
	Expand      interface{} `json:"expand" yaml:"expand"` // 扩展配置
}

type Logger struct {
	Level  string `json:"level" yaml:"level"`
	OutDir string `json:"out_dir" yaml:"out-dir"`
	Format string `json:"format" yaml:"format"`
}

// GetDefaultConfig get default web app configuration
func getDefaultConfig() *Application {
	return &Application{
		Host:        "0.0.0.0",
		Port:        8080,
		ContextPath: "/",
		Logger:      Logger{Level: "debug", OutDir: "./log", Format: "json"},
	}
}

// ReadApplicationConfigurationFile 读取配置文件
func ReadApplicationConfigurationFile(configPath string) {
	if configPath == "" {
		//configPath = "./config/application.yml"
		Runtime.Application = getDefaultConfig()
		InitRuntime()
		Runtime.Log.Debug("loading default application configuration")
		return
	}
	if dbs, err := ioutil.ReadFile(configPath); err == nil {
		Runtime.Application = new(Application)
		_ = yaml.Unmarshal(dbs, Runtime.Application)
	}
	InitRuntime()
}

// GetHost get host string
func (a *Application) GetHost() string {
	host := bytes.Buffer{}
	host.WriteString(a.Host)
	host.WriteString(":")
	host.WriteString(strconv.Itoa(a.Port))
	return host.String()
}
