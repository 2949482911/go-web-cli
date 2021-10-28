package entrance

import (
	"github.com/go-web-cli/config"
	"github.com/go-web-cli/router"
)

var (
	configPath = ""
)

func init() {

}

func StartApp() {
	//  1.读取配置文件
	config.ReadApplicationConfigurationFile(configPath)
	//  2.初始化运行时对象
	config.InitRuntime()
	config.Runtime.Log.Info("init runtime success")
	//  2.加载路由
	router.ExeRouter()
	//  4.启动
	config.StartEngine()
}
