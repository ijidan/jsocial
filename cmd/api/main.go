package main

import (
	"context"
	"flag"
	"github.com/ijidan/jsocial/internal/app/api/router"
	"github.com/ijidan/jsocial/internal/global"
	"github.com/ijidan/jsocial/internal/injector"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/ijidan/jsocial/internal/pkg/file"
	"github.com/ijidan/jsocial/internal/pkg/info"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var path = file.NewPath()
var configFile = flag.String("c", path.ConfigsDir+"api.yml", "set config file which viper will loading.")
var rootPath = flag.String("d", path.ProjectDir, "set project path")

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	param := config.NewCmdParam(*rootPath, *configFile)
	global.GH, _ = injector.InitializeEventHttpGlobal(*param)
	r := router.NewGin(global.GH.Conf.App)
	// Send output
	table := info.BuildApiTable(global.GH.Conf.Http, os.Getpid())
	table.Render()
	addr := info.BuildApiAddr(global.GH.Conf.Http)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")

}
