package cmd

import (
	"fmt"
	"github.com/dotdancer/gogofly/conf"
	"github.com/dotdancer/gogofly/router"
)

func Start() {
	conf.InitConfig()
	router.InitRouter()
}

func Clean() {
	fmt.Println("========Clean==========")
}
