package log

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
)

type LoggerPlugin struct {
}

func (lp *LoggerPlugin) Name() string {
	return "sql_log"
}

func (lp *LoggerPlugin) Initialize(db *gorm.DB) error {
	beforeMap := map[string]func(*gorm.DB){
		"gorm:create": writeLogBefore,
		"gorm:delete": writeLogBefore,
		"gorm:update": writeLogBefore,
		"gorm:query":  writeLogBefore,
		"gorm:row":    writeLogBefore,
		"gorm:raw":    writeLogBefore,
	}
	for k, v := range beforeMap {
		_ = db.Callback().Create().Before(k).Register(lp.Name()+k+"_before", v)
	}
	afterMap := map[string]func(db *gorm.DB){
		"gorm:create": writeLogAfter,
		"gorm:delete": writeLogAfter,
		"gorm:update": writeLogAfter,
		"gorm:query":  writeLogAfter,
		"gorm:row":    writeLogAfter,
		"gorm:raw":    writeLogAfter,
	}
	for k, v := range afterMap {
		_ = db.Callback().Create().After(k).Register(lp.Name()+k+"_after", v)
	}

	return nil
}

func writeLogBefore(db *gorm.DB) {
	db.InstanceSet("start_time", carbon.Now().TimestampMilli())
}

func writeLogAfter(db *gorm.DB) {
	endTime := carbon.Now().TimestampMilli()
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	startTime, _ := db.InstanceGet("start_time")
	log := fmt.Sprintf("start time:%d,end time:%d,sql:%s", startTime, endTime, sql)
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
	logrus.Error(log)
}
