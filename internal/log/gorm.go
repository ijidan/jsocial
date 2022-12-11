package log

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type LoggerPlugin struct {
}

func (lp *LoggerPlugin) Name() string {
	return "sql_log"
}

func (lp *LoggerPlugin) Initialize(db *gorm.DB) error {

	_ = db.Callback().Create().Before("gorm:create").Register(lp.Name()+":gorm:create:before", writeLogBefore)
	_ = db.Callback().Delete().Before("gorm:delete").Register(lp.Name()+":gorm:delete:before", writeLogBefore)
	_ = db.Callback().Update().Before("gorm:update").Register(lp.Name()+":gorm:update:before", writeLogBefore)
	_ = db.Callback().Query().Before("gorm:query").Register(lp.Name()+":gorm:query:before", writeLogBefore)
	_ = db.Callback().Row().Before("gorm:row").Register(lp.Name()+":gorm:row:before", writeLogBefore)
	_ = db.Callback().Raw().Before("gorm:raw").Register(lp.Name()+":gorm:raw:before", writeLogBefore)

	_ = db.Callback().Create().After("gorm:create").Register(lp.Name()+":gorm:create:after", writeLogAfter)
	_ = db.Callback().Delete().After("gorm:delete").Register(lp.Name()+":gorm:delete:after", writeLogAfter)
	_ = db.Callback().Update().After("gorm:update").Register(lp.Name()+":gorm:update:after", writeLogAfter)
	_ = db.Callback().Query().After("gorm:query").Register(lp.Name()+":gorm:query:after", writeLogAfter)
	_ = db.Callback().Row().After("gorm:row").Register(lp.Name()+":gorm:row:after", writeLogAfter)
	_ = db.Callback().Raw().After("gorm:raw").Register(lp.Name()+":gorm:raw:after", writeLogAfter)

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
	color.Green(log)
}
