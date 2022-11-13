package logger

import (
	"github.com/google/wire"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"io"
	"os"
)

type Root string

func NewLogger(conf *config.Http, root Root) *logrus.Logger {
	logfile := cast.ToString(root) + "/" + conf.Log
	writer1 := os.Stdout
	writer2, _ := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE, 0755)
	instanceLogger := &logrus.Logger{}
	instanceLogger.SetLevel(logrus.InfoLevel)
	instanceLogger.SetFormatter(&logrus.JSONFormatter{})
	instanceLogger.SetOutput(io.MultiWriter(writer1, writer2))
	return instanceLogger
}

var Provider = wire.NewSet(NewLogger)
