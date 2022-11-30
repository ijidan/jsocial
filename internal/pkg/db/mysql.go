package db

import (
	"fmt"
	"github.com/google/wire"
	"github.com/ijidan/jsocial/internal/log"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func NewDb(conf *config.Mysql) *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.Database, conf.Charset)
	instanceDb, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	sqlDB, _ := instanceDb.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	instanceDb.Use(&log.LoggerPlugin{})
	return instanceDb
}

var Provider = wire.NewSet(NewDb)
