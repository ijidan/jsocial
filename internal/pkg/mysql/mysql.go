package mysql

import (
	"fmt"
	"github.com/google/wire"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
	"time"
)

var (
	onceDb     sync.Once
	instanceDb *gorm.DB
)

func GetDbInstance(cf config.Mysql) (*gorm.DB,error) {
	onceDb.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", cf.User, cf.Password, cf.Host, cf.Port, cf.Database, cf.Charset)
		instanceDb, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			//Logger: logger.Default.LogMode(logger.Info),
		})
		sqlDB, _ := instanceDb.DB()
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(10)
		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(100)
		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour)

	})

	return instanceDb,nil
}

var Provider=wire.NewSet(GetDbInstance)
