package global

import (
	"github.com/gomodule/redigo/redis"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/ijidan/jsocial/internal/pkg/jaeger"
	"github.com/ijidan/jsocial/internal/pkg/logger"
	"github.com/ijidan/jsocial/internal/pkg/response"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
	"path/filepath"
	"runtime"
)

var (
	Root      string
	Conf      *config.Config
	Logger    *logrus.Logger
	Db        *gorm.DB
	Rd        redis.Conn
	Rsp       *response.Response
	Tracer    opentracing.Tracer
	Closer    io.Closer
	RequestId string
)

func Close() {
	//sqlDB, _ := Db.DB()
	//_ = sqlDB.Close()
	//_ = Rd.Close()
	_ = Closer.Close()
}
func init() {
	_, file, _, _ := runtime.Caller(0)
	Root = filepath.Dir(filepath.Dir(file))
	Conf = config.GetConfigInstance(Root)
	Logger, _ = logger.GetLoggerInstance(Conf, Root)
	//Db = GetDbInstance(Conf)
	//Rd = GetRdInstance(Conf)
	Rsp, _ = response.GetResponseInstance()
	Tracer, Closer, _ = jaeger.NewJaeger(Conf, "jim_service")
	RequestId = "X-Request-Id"
}
