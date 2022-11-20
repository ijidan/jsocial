package global

import (
	"github.com/gomodule/redigo/redis"
	"github.com/google/wire"
	"github.com/ijidan/jsocial/internal/constant"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/ijidan/jsocial/internal/pkg/jaeger"
	"github.com/ijidan/jsocial/internal/pkg/response"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
)

type RpcGlobal struct {
	Conf      *config.Config
	Logger    *logrus.Logger
	Db        *gorm.DB
	Rd        redis.Conn
	Rsp       *response.Response
	Tracer    opentracing.Tracer
	Closer    io.Closer
	RequestId string
}

func NewRpcGlobal(conf *config.Config, logger *logrus.Logger, db *gorm.DB, rd redis.Conn, rsp *response.Response, j *jaeger.Jaeger) *RpcGlobal {
	instanceGlobal := &RpcGlobal{
		Conf:      conf,
		Logger:    logger,
		Db:        db,
		Rd:        rd,
		Rsp:       rsp,
		Tracer:    j.Tracer,
		Closer:    j.Closer,
		RequestId: constant.RequestId,
	}

	return instanceGlobal
}

var GR *RpcGlobal

var RpcProvider = wire.NewSet(NewRpcGlobal)
