package entrance

import (
	"flag"
	"github.com/2949482911/go-web-cli/config"
	"github.com/2949482911/go-web-cli/router"
)

var (
	configPath = ""
)

func init() {
	flag.StringVar(&configPath, "c", "", "配置文件地址默认为当前项目下,config/application.yml")
	flag.Parse()
}

func StartApp() {
	//  1.读取配置文件
	config.ReadApplicationConfigurationFile(configPath)
	//  3.加载路由
	router.ExeRouter()
	//  4.启动
	config.StartEngine()
}
