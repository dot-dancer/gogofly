package cmd

import (
	"fmt"
	"github.com/dotdancer/gogofly/conf"
	"github.com/dotdancer/gogofly/global"
	"github.com/dotdancer/gogofly/router"
)

func Start() {
	// ===============================================================================
	// = 初始化系统配置文件
	conf.InitConfig()

	// ===============================================================================
	// = 初始化日志组件
	global.Logger = conf.InitLogger()

	// ===============================================================================
	// = 初始化系统路由
	router.InitRouter()
}

func Clean() {
	fmt.Println("========Clean==========")
}
