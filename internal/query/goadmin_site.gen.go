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

func newGoadminSite(db *gorm.DB) goadminSite {
	_goadminSite := goadminSite{}

	_goadminSite.goadminSiteDo.UseDB(db)
	_goadminSite.goadminSiteDo.UseModel(&model.GoadminSite{})

	tableName := _goadminSite.goadminSiteDo.TableName()
	_goadminSite.ALL = field.NewAsterisk(tableName)
	_goadminSite.ID = field.NewInt32(tableName, "id")
	_goadminSite.Key = field.NewString(tableName, "key")
	_goadminSite.Value = field.NewString(tableName, "value")
	_goadminSite.Description = field.NewString(tableName, "description")
	_goadminSite.State = field.NewInt32(tableName, "state")
	_goadminSite.CreatedAt = field.NewTime(tableName, "created_at")
	_goadminSite.UpdatedAt = field.NewTime(tableName, "updated_at")

	_goadminSite.fillFieldMap()

	return _goadminSite
}

type goadminSite struct {
	goadminSiteDo goadminSiteDo

	ALL         field.Asterisk
	ID          field.Int32
	Key         field.String
	Value       field.String
	Description field.String
	State       field.Int32
	CreatedAt   field.Time
	UpdatedAt   field.Time

	fieldMap map[string]field.Expr
}

func (g goadminSite) Table(newTableName string) *goadminSite {
	g.goadminSiteDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g goadminSite) As(alias string) *goadminSite {
	g.goadminSiteDo.DO = *(g.goadminSiteDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *goadminSite) updateTableName(table string) *goadminSite {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt32(table, "id")
	g.Key = field.NewString(table, "key")
	g.Value = field.NewString(table, "value")
	g.Description = field.NewString(table, "description")
	g.State = field.NewInt32(table, "state")
	g.CreatedAt = field.NewTime(table, "created_at")
	g.UpdatedAt = field.NewTime(table, "updated_at")

	g.fillFieldMap()

	return g
}

func (g *goadminSite) WithContext(ctx context.Context) IGoadminSiteDo {
	return g.goadminSiteDo.WithContext(ctx)
}

func (g goadminSite) TableName() string { return g.goadminSiteDo.TableName() }

func (g goadminSite) Alias() string { return g.goadminSiteDo.Alias() }

func (g *goadminSite) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *goadminSite) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 7)
	g.fieldMap["id"] = g.ID
	g.fieldMap["key"] = g.Key
	g.fieldMap["value"] = g.Value
	g.fieldMap["description"] = g.Description
	g.fieldMap["state"] = g.State
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
}

func (g goadminSite) clone(db *gorm.DB) goadminSite {
	g.goadminSiteDo.ReplaceDB(db)
	return g
}

type goadminSiteDo struct{ gen.DO }

type IGoadminSiteDo interface {
	gen.SubQuery
	Debug() IGoadminSiteDo
	WithContext(ctx context.Context) IGoadminSiteDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGoadminSiteDo
	WriteDB() IGoadminSiteDo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGoadminSiteDo
	Not(conds ...gen.Condition) IGoadminSiteDo
	Or(conds ...gen.Condition) IGoadminSiteDo
	Select(conds ...field.Expr) IGoadminSiteDo
	Where(conds ...gen.Condition) IGoadminSiteDo
	Order(conds ...field.Expr) IGoadminSiteDo
	Distinct(cols ...field.Expr) IGoadminSiteDo
	Omit(cols ...field.Expr) IGoadminSiteDo
	Join(table schema.Tabler, on ...field.Expr) IGoadminSiteDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGoadminSiteDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGoadminSiteDo
	Group(cols ...field.Expr) IGoadminSiteDo
	Having(conds ...gen.Condition) IGoadminSiteDo
	Limit(limit int) IGoadminSiteDo
	Offset(offset int) IGoadminSiteDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGoadminSiteDo
	Unscoped() IGoadminSiteDo
	Create(values ...*model.GoadminSite) error
	CreateInBatches(values []*model.GoadminSite, batchSize int) error
	Save(values ...*model.GoadminSite) error
	First() (*model.GoadminSite, error)
	Take() (*model.GoadminSite, error)
	Last() (*model.GoadminSite, error)
	Find() ([]*model.GoadminSite, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GoadminSite, err error)
	FindInBatches(result *[]*model.GoadminSite, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.GoadminSite) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGoadminSiteDo
	Assign(attrs ...field.AssignExpr) IGoadminSiteDo
	Joins(fields ...field.RelationField) IGoadminSiteDo
	Preload(fields ...field.RelationField) IGoadminSiteDo
	FirstOrInit() (*model.GoadminSite, error)
	FirstOrCreate() (*model.GoadminSite, error)
	FindByPage(offset int, limit int) (result []*model.GoadminSite, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGoadminSiteDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetById(id int) (result *model.GoadminSite, err error)
}

//SELECT * FROM @@table WHERE id=@id
func (g goadminSiteDo) GetById(id int) (result *model.GoadminSite, err error) {
	params := make(map[string]interface{}, 0)

	var generateSQL strings.Builder
	params["id"] = id
	generateSQL.WriteString("SELECT * FROM goadmin_site WHERE id=@id ")

	var executeSQL *gorm.DB
	if len(params) > 0 {
		executeSQL = g.UnderlyingDB().Raw(generateSQL.String(), params).Take(&result)
	} else {
		executeSQL = g.UnderlyingDB().Raw(generateSQL.String()).Take(&result)
	}
	err = executeSQL.Error
	return
}

func (g goadminSiteDo) Debug() IGoadminSiteDo {
	return g.withDO(g.DO.Debug())
}

func (g goadminSiteDo) WithContext(ctx context.Context) IGoadminSiteDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g goadminSiteDo) ReadDB() IGoadminSiteDo {
	return g.Clauses(dbresolver.Read)
}

func (g goadminSiteDo) WriteDB() IGoadminSiteDo {
	return g.Clauses(dbresolver.Write)
}

func (g goadminSiteDo) Clauses(conds ...clause.Expression) IGoadminSiteDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g goadminSiteDo) Returning(value interface{}, columns ...string) IGoadminSiteDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g goadminSiteDo) Not(conds ...gen.Condition) IGoadminSiteDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g goadminSiteDo) Or(conds ...gen.Condition) IGoadminSiteDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g goadminSiteDo) Select(conds ...field.Expr) IGoadminSiteDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g goadminSiteDo) Where(conds ...gen.Condition) IGoadminSiteDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g goadminSiteDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IGoadminSiteDo {
	return g.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (g goadminSiteDo) Order(conds ...field.Expr) IGoadminSiteDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g goadminSiteDo) Distinct(cols ...field.Expr) IGoadminSiteDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g goadminSiteDo) Omit(cols ...field.Expr) IGoadminSiteDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g goadminSiteDo) Join(table schema.Tabler, on ...field.Expr) IGoadminSiteDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g goadminSiteDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGoadminSiteDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g goadminSiteDo) RightJoin(table schema.Tabler, on ...field.Expr) IGoadminSiteDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g goadminSiteDo) Group(cols ...field.Expr) IGoadminSiteDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g goadminSiteDo) Having(conds ...gen.Condition) IGoadminSiteDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g goadminSiteDo) Limit(limit int) IGoadminSiteDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g goadminSiteDo) Offset(offset int) IGoadminSiteDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g goadminSiteDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGoadminSiteDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g goadminSiteDo) Unscoped() IGoadminSiteDo {
	return g.withDO(g.DO.Unscoped())
}

func (g goadminSiteDo) Create(values ...*model.GoadminSite) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g goadminSiteDo) CreateInBatches(values []*model.GoadminSite, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g goadminSiteDo) Save(values ...*model.GoadminSite) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g goadminSiteDo) First() (*model.GoadminSite, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminSite), nil
	}
}

func (g goadminSiteDo) Take() (*model.GoadminSite, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminSite), nil
	}
}

func (g goadminSiteDo) Last() (*model.GoadminSite, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminSite), nil
	}
}

func (g goadminSiteDo) Find() ([]*model.GoadminSite, error) {
	result, err := g.DO.Find()
	return result.([]*model.GoadminSite), err
}

func (g goadminSiteDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GoadminSite, err error) {
	buf := make([]*model.GoadminSite, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g goadminSiteDo) FindInBatches(result *[]*model.GoadminSite, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g goadminSiteDo) Attrs(attrs ...field.AssignExpr) IGoadminSiteDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g goadminSiteDo) Assign(attrs ...field.AssignExpr) IGoadminSiteDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g goadminSiteDo) Joins(fields ...field.RelationField) IGoadminSiteDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g goadminSiteDo) Preload(fields ...field.RelationField) IGoadminSiteDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g goadminSiteDo) FirstOrInit() (*model.GoadminSite, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminSite), nil
	}
}

func (g goadminSiteDo) FirstOrCreate() (*model.GoadminSite, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminSite), nil
	}
}

func (g goadminSiteDo) FindByPage(offset int, limit int) (result []*model.GoadminSite, count int64, err error) {
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

func (g goadminSiteDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g goadminSiteDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g goadminSiteDo) Delete(models ...*model.GoadminSite) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *goadminSiteDo) withDO(do gen.Dao) *goadminSiteDo {
	g.DO = *do.(*gen.DO)
	return g
}
