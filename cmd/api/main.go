package main

import (
	"flag"
	"github.com/ijidan/jsocial/internal/app/api/router"
	"github.com/ijidan/jsocial/internal/global"
	"github.com/ijidan/jsocial/internal/injector"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/ijidan/jsocial/internal/pkg/info"
	"os"
)

var configFile = flag.String("f", "E:\\go_project\\jsocial\\configs\\api.yml", "set config file which viper will loading.")
var rootPath = flag.String("r", "E:\\go_project\\jsocial", "set root path")

func main() {
	param := config.NewCmdParam(*rootPath, *configFile)
	global.GH, _ = injector.InitializeEventHttpGlobal(*param)
	r := router.NewGin(global.GH.Conf.App)
	// Send output
	table := info.BuildApiTable(global.GH.Conf.Http, os.Getpid())
	table.Render()
	addr := info.BuildApiAddr(global.GH.Conf.Http)
	if err := r.Run(addr); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
