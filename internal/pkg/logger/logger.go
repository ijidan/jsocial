package logger

import (
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"sync"
)

var (
	onceLogger     sync.Once
	instanceLogger *logrus.Logger
)

func GetLoggerInstance(config *config.Config, root string) (*logrus.Logger,error) {
	onceLogger.Do(func() {
		logfile := root + "/" + config.Rpc.Log
		writer1 := os.Stdout
		writer2, _ := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE, 0755)
		instanceLogger = logrus.New()
		instanceLogger.SetLevel(logrus.InfoLevel)
		instanceLogger.SetFormatter(&logrus.JSONFormatter{})
		instanceLogger.SetOutput(io.MultiWriter(writer1, writer2))
	})
	return instanceLogger,nil
}
