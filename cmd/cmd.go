package cmd

import (
	"fmt"
	"github.com/dotdancer/gogofly/conf"
	"github.com/dotdancer/gogofly/global"
	"github.com/dotdancer/gogofly/router"
	"github.com/dotdancer/gogofly/utils"
)

func Start() {
	var initErr error

	// ===============================================================================
	// = 初始化系统配置文件
	conf.InitConfig()

	// ===============================================================================
	// = 初始化日志组件
	global.Logger = conf.InitLogger()

	// ===============================================================================
	// = 初始化数据库连接
	db, err := conf.InitDB()
	global.DB = db
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}

	// ===============================================================================
	// = 初始化Redis连接
	rdClient, err := conf.InitRedis()
	global.RedisClient = rdClient
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}

	// ===============================================================================
	// = 初始化过程中, 遇到错误的最终处理
	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}

		panic(initErr.Error())
	}

	// ===============================================================================
	// = 初始化系统路由
	router.InitRouter()
}

func Clean() {
	fmt.Println("========Clean==========")
}
