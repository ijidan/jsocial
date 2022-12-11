package main

import (
	"context"
	"flag"
	"github.com/fatih/color"
	"github.com/ijidan/jsocial/internal/constant"
	"github.com/ijidan/jsocial/internal/global"
	"github.com/ijidan/jsocial/internal/injector"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/ijidan/jsocial/internal/pkg/file"
	"github.com/ijidan/jsocial/internal/pkg/info"
	"github.com/ijidan/jsocial/internal/server"
	"os"
	"os/signal"
	"syscall"
)

var path = file.NewPath()
var configFile = flag.String("c", path.ConfigsDir+"gateway.yml", "set config file which viper will loading.")
var rootPath = flag.String("d", path.ProjectDir, "set project path")

func main() {
	param := config.NewCmdParam(*rootPath, *configFile)
	global.GR, _ = injector.InitializeEventServiceGlobal(*param)
	table := info.BuildRpcTable(*global.GR.Conf)
	table.Render()
	ctx, cancel := context.WithCancel(context.Background())
	server.StartRpc(ctx, cancel, constant.ServiceGateway)

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-ch
	color.Red("closing ...")
	cancel()
	color.Red("closed")
}
