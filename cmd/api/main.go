package main

import (
	"flag"
	"github.com/fvbock/endless"
	"github.com/ijidan/jsocial/internal/app/api/router"
	"github.com/ijidan/jsocial/internal/global"
	"github.com/ijidan/jsocial/internal/injector"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/ijidan/jsocial/internal/pkg/file"
	"github.com/ijidan/jsocial/internal/pkg/info"
	"log"
	"os"
)

var path = file.NewPath()
var configFile = flag.String("c", path.ConfigsDir+"api.yml", "set config file which viper will loading.")
var rootPath = flag.String("d", path.ProjectDir, "set project path")

func main() {
	param := config.NewCmdParam(*rootPath, *configFile)
	global.GH, _ = injector.InitializeEventHttpGlobal(*param)
	r := router.NewGin(global.GH.Conf.App)
	// Send output
	table := info.BuildApiTable(global.GH.Conf.Http, os.Getpid())
	table.Render()
	addr := info.BuildApiAddr(global.GH.Conf.Http)
	if err := endless.ListenAndServe(addr, r); err != nil {
		log.Fatalf("listen: %s\n", err)
	}
}
