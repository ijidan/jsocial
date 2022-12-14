// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/ijidan/jsocial/internal/model"
)

func newGoadminOperationLog(db *gorm.DB) goadminOperationLog {
	_goadminOperationLog := goadminOperationLog{}

	_goadminOperationLog.goadminOperationLogDo.UseDB(db)
	_goadminOperationLog.goadminOperationLogDo.UseModel(&model.GoadminOperationLog{})

	tableName := _goadminOperationLog.goadminOperationLogDo.TableName()
	_goadminOperationLog.ALL = field.NewAsterisk(tableName)
	_goadminOperationLog.ID = field.NewInt32(tableName, "id")
	_goadminOperationLog.UserID = field.NewInt32(tableName, "user_id")
	_goadminOperationLog.Path = field.NewString(tableName, "path")
	_goadminOperationLog.Method = field.NewString(tableName, "method")
	_goadminOperationLog.IP = field.NewString(tableName, "ip")
	_goadminOperationLog.Input = field.NewString(tableName, "input")
	_goadminOperationLog.CreatedAt = field.NewTime(tableName, "created_at")
	_goadminOperationLog.UpdatedAt = field.NewTime(tableName, "updated_at")

	_goadminOperationLog.fillFieldMap()

	return _goadminOperationLog
}

type goadminOperationLog struct {
	goadminOperationLogDo goadminOperationLogDo

	ALL       field.Asterisk
	ID        field.Int32
	UserID    field.Int32
	Path      field.String
	Method    field.String
	IP        field.String
	Input     field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (g goadminOperationLog) Table(newTableName string) *goadminOperationLog {
	g.goadminOperationLogDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g goadminOperationLog) As(alias string) *goadminOperationLog {
	g.goadminOperationLogDo.DO = *(g.goadminOperationLogDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *goadminOperationLog) updateTableName(table string) *goadminOperationLog {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt32(table, "id")
	g.UserID = field.NewInt32(table, "user_id")
	g.Path = field.NewString(table, "path")
	g.Method = field.NewString(table, "method")
	g.IP = field.NewString(table, "ip")
	g.Input = field.NewString(table, "input")
	g.CreatedAt = field.NewTime(table, "created_at")
	g.UpdatedAt = field.NewTime(table, "updated_at")

	g.fillFieldMap()

	return g
}

func (g *goadminOperationLog) WithContext(ctx context.Context) IGoadminOperationLogDo {
	return g.goadminOperationLogDo.WithContext(ctx)
}

func (g goadminOperationLog) TableName() string { return g.goadminOperationLogDo.TableName() }

func (g goadminOperationLog) Alias() string { return g.goadminOperationLogDo.Alias() }

func (g *goadminOperationLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *goadminOperationLog) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 8)
	g.fieldMap["id"] = g.ID
	g.fieldMap["user_id"] = g.UserID
	g.fieldMap["path"] = g.Path
	g.fieldMap["method"] = g.Method
	g.fieldMap["ip"] = g.IP
	g.fieldMap["input"] = g.Input
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
}

func (g goadminOperationLog) clone(db *gorm.DB) goadminOperationLog {
	g.goadminOperationLogDo.ReplaceDB(db)
	return g
}

type goadminOperationLogDo struct{ gen.DO }

type IGoadminOperationLogDo interface {
	gen.SubQuery
	Debug() IGoadminOperationLogDo
	WithContext(ctx context.Context) IGoadminOperationLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGoadminOperationLogDo
	WriteDB() IGoadminOperationLogDo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGoadminOperationLogDo
	Not(conds ...gen.Condition) IGoadminOperationLogDo
	Or(conds ...gen.Condition) IGoadminOperationLogDo
	Select(conds ...field.Expr) IGoadminOperationLogDo
	Where(conds ...gen.Condition) IGoadminOperationLogDo
	Order(conds ...field.Expr) IGoadminOperationLogDo
	Distinct(cols ...field.Expr) IGoadminOperationLogDo
	Omit(cols ...field.Expr) IGoadminOperationLogDo
	Join(table schema.Tabler, on ...field.Expr) IGoadminOperationLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGoadminOperationLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGoadminOperationLogDo
	Group(cols ...field.Expr) IGoadminOperationLogDo
	Having(conds ...gen.Condition) IGoadminOperationLogDo
	Limit(limit int) IGoadminOperationLogDo
	Offset(offset int) IGoadminOperationLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGoadminOperationLogDo
	Unscoped() IGoadminOperationLogDo
	Create(values ...*model.GoadminOperationLog) error
	CreateInBatches(values []*model.GoadminOperationLog, batchSize int) error
	Save(values ...*model.GoadminOperationLog) error
	First() (*model.GoadminOperationLog, error)
	Take() (*model.GoadminOperationLog, error)
	Last() (*model.GoadminOperationLog, error)
	Find() ([]*model.GoadminOperationLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GoadminOperationLog, err error)
	FindInBatches(result *[]*model.GoadminOperationLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.GoadminOperationLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGoadminOperationLogDo
	Assign(attrs ...field.AssignExpr) IGoadminOperationLogDo
	Joins(fields ...field.RelationField) IGoadminOperationLogDo
	Preload(fields ...field.RelationField) IGoadminOperationLogDo
	FirstOrInit() (*model.GoadminOperationLog, error)
	FirstOrCreate() (*model.GoadminOperationLog, error)
	FindByPage(offset int, limit int) (result []*model.GoadminOperationLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGoadminOperationLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetById(id int) (result *model.GoadminOperationLog, err error)
}

//SELECT * FROM @@table WHERE id=@id
func (g goadminOperationLogDo) GetById(id int) (result *model.GoadminOperationLog, err error) {
	params := make(map[string]interface{}, 0)

	var generateSQL strings.Builder
	params["id"] = id
	generateSQL.WriteString("SELECT * FROM goadmin_operation_log WHERE id=@id ")

	var executeSQL *gorm.DB
	if len(params) > 0 {
		executeSQL = g.UnderlyingDB().Raw(generateSQL.String(), params).Take(&result)
	} else {
		executeSQL = g.UnderlyingDB().Raw(generateSQL.String()).Take(&result)
	}
	err = executeSQL.Error
	return
}

func (g goadminOperationLogDo) Debug() IGoadminOperationLogDo {
	return g.withDO(g.DO.Debug())
}

func (g goadminOperationLogDo) WithContext(ctx context.Context) IGoadminOperationLogDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g goadminOperationLogDo) ReadDB() IGoadminOperationLogDo {
	return g.Clauses(dbresolver.Read)
}

func (g goadminOperationLogDo) WriteDB() IGoadminOperationLogDo {
	return g.Clauses(dbresolver.Write)
}

func (g goadminOperationLogDo) Clauses(conds ...clause.Expression) IGoadminOperationLogDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g goadminOperationLogDo) Returning(value interface{}, columns ...string) IGoadminOperationLogDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g goadminOperationLogDo) Not(conds ...gen.Condition) IGoadminOperationLogDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g goadminOperationLogDo) Or(conds ...gen.Condition) IGoadminOperationLogDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g goadminOperationLogDo) Select(conds ...field.Expr) IGoadminOperationLogDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g goadminOperationLogDo) Where(conds ...gen.Condition) IGoadminOperationLogDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g goadminOperationLogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IGoadminOperationLogDo {
	return g.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (g goadminOperationLogDo) Order(conds ...field.Expr) IGoadminOperationLogDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g goadminOperationLogDo) Distinct(cols ...field.Expr) IGoadminOperationLogDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g goadminOperationLogDo) Omit(cols ...field.Expr) IGoadminOperationLogDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g goadminOperationLogDo) Join(table schema.Tabler, on ...field.Expr) IGoadminOperationLogDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g goadminOperationLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGoadminOperationLogDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g goadminOperationLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IGoadminOperationLogDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g goadminOperationLogDo) Group(cols ...field.Expr) IGoadminOperationLogDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g goadminOperationLogDo) Having(conds ...gen.Condition) IGoadminOperationLogDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g goadminOperationLogDo) Limit(limit int) IGoadminOperationLogDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g goadminOperationLogDo) Offset(offset int) IGoadminOperationLogDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g goadminOperationLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGoadminOperationLogDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g goadminOperationLogDo) Unscoped() IGoadminOperationLogDo {
	return g.withDO(g.DO.Unscoped())
}

func (g goadminOperationLogDo) Create(values ...*model.GoadminOperationLog) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g goadminOperationLogDo) CreateInBatches(values []*model.GoadminOperationLog, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g goadminOperationLogDo) Save(values ...*model.GoadminOperationLog) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g goadminOperationLogDo) First() (*model.GoadminOperationLog, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminOperationLog), nil
	}
}

func (g goadminOperationLogDo) Take() (*model.GoadminOperationLog, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminOperationLog), nil
	}
}

func (g goadminOperationLogDo) Last() (*model.GoadminOperationLog, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminOperationLog), nil
	}
}

func (g goadminOperationLogDo) Find() ([]*model.GoadminOperationLog, error) {
	result, err := g.DO.Find()
	return result.([]*model.GoadminOperationLog), err
}

func (g goadminOperationLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GoadminOperationLog, err error) {
	buf := make([]*model.GoadminOperationLog, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g goadminOperationLogDo) FindInBatches(result *[]*model.GoadminOperationLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g goadminOperationLogDo) Attrs(attrs ...field.AssignExpr) IGoadminOperationLogDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g goadminOperationLogDo) Assign(attrs ...field.AssignExpr) IGoadminOperationLogDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g goadminOperationLogDo) Joins(fields ...field.RelationField) IGoadminOperationLogDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g goadminOperationLogDo) Preload(fields ...field.RelationField) IGoadminOperationLogDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g goadminOperationLogDo) FirstOrInit() (*model.GoadminOperationLog, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminOperationLog), nil
	}
}

func (g goadminOperationLogDo) FirstOrCreate() (*model.GoadminOperationLog, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminOperationLog), nil
	}
}

func (g goadminOperationLogDo) FindByPage(offset int, limit int) (result []*model.GoadminOperationLog, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g goadminOperationLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g goadminOperationLogDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g goadminOperationLogDo) Delete(models ...*model.GoadminOperationLog) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *goadminOperationLogDo) withDO(do gen.Dao) *goadminOperationLogDo {
	g.DO = *do.(*gen.DO)
	return g
}
