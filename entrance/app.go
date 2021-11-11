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

func StartApp(expand interface{}, fuc ...func()) {

	if configPath == "" {
		print("configPath is empty ")
		return
	}
	//  1.读取配置文件
	config.ReadApplicationConfigurationFile(configPath, expand)
	//  3.加载路由
	router.ExeRouter()

	config.ExternalFunc(fuc...)
	//  4.启动
	config.StartEngine()
}
