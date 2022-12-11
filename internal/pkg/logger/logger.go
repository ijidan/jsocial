package logger

import (
	"github.com/google/wire"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

func NewLogger(conf *config.Http, param config.CmdParam) *logrus.Logger {
	logfile := param.RootPath + "/" + conf.Log
	writer1 := os.Stdout
	writer2, _ := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE, 0755)
	instanceLogger := &logrus.Logger{}
	instanceLogger.SetLevel(logrus.InfoLevel)
	instanceLogger.SetFormatter(&logrus.JSONFormatter{})
	instanceLogger.SetOutput(io.MultiWriter(writer1, writer2))
	return instanceLogger
}

var Provider = wire.NewSet(NewLogger)

func WriteLog(log string) {
	var logger = &lumberjack.Logger{
		Filename:   "./log.txt",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days

	}
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006/01/02 - 15:04:05",
	})
	//同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(os.Stdout, logger)
	logrus.SetOutput(fileAndStdoutWriter)
	//设置最低loglevel
	logrus.SetLevel(logrus.InfoLevel)
	logrus.Info(log)
}
