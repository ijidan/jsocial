//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"github.com/ijidan/jsocial/internal/pkg/db"
	"github.com/ijidan/jsocial/internal/pkg/global"
	"github.com/ijidan/jsocial/internal/pkg/jaeger"
	"github.com/ijidan/jsocial/internal/pkg/logger"
	"github.com/ijidan/jsocial/internal/pkg/redis"
	"github.com/ijidan/jsocial/internal/pkg/response"
)

func InitializeEventHttpGlobal(root logger.Root, configPath string) (*global.HttpGlobal, error) {
	wire.Build(
		config.AppProvider,
		config.MysqlProvider,
		config.HttpProvider,
		config.JaegerProvider,
		config.RedisProvider,
		config.Provider,
		global.HttpProvider,
		logger.Provider,
		db.Provider,
		redis.Provider,
		response.Provider,
		jaeger.Provider)
	return &global.HttpGlobal{}, nil
}
