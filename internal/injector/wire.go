//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"github.com/ijidan/jsocial/internal/global"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/ijidan/jsocial/internal/pkg/db"
	"github.com/ijidan/jsocial/internal/pkg/jaeger"
	"github.com/ijidan/jsocial/internal/pkg/logger"
	"github.com/ijidan/jsocial/internal/pkg/redis"
	"github.com/ijidan/jsocial/internal/pkg/response"
)

func InitializeEventHttpGlobal(param config.CmdParam) (*global.HttpGlobal, error) {
	wire.Build(
		config.AppProvider,
		config.MysqlProvider,
		config.JaegerProvider,
		config.RedisProvider,
		config.HttpProvider,

		global.HttpProvider,
		config.Provider,
		logger.Provider,
		db.Provider,
		redis.Provider,
		response.Provider,
		jaeger.Provider)
	return &global.HttpGlobal{}, nil
}

func InitializeEventServiceGlobal(param config.CmdParam) (*global.RpcGlobal, error) {
	wire.Build(
		config.AppProvider,
		config.MysqlProvider,
		config.JaegerProvider,
		config.RedisProvider,
		config.HttpProvider,

		global.RpcProvider,
		config.Provider,
		logger.Provider,
		db.Provider,
		redis.Provider,
		response.Provider,
		jaeger.Provider)
	return &global.RpcGlobal{}, nil
}
