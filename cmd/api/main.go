package main

import (
	"flag"
	"github.com/ijidan/jsocial/internal/app/api/router"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/ijidan/jsocial/internal/pkg/info"
	"os"
)

var Config *config.Config

var configFile = flag.String("f", "api.yml", "set config file which viper will loading.")

func main() {
	Config = config.GetConfigInstance(*configFile)
	r := router.NewGin(Config.App)
	// Send output
	table := info.BuildApiTable(Config.Http, os.Getpid())
	table.Render()
	addr := info.BuildApiAddr(Config.Http)
	if err := r.Run(addr); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
