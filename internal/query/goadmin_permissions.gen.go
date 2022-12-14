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

func newGoadminPermission(db *gorm.DB) goadminPermission {
	_goadminPermission := goadminPermission{}

	_goadminPermission.goadminPermissionDo.UseDB(db)
	_goadminPermission.goadminPermissionDo.UseModel(&model.GoadminPermission{})

	tableName := _goadminPermission.goadminPermissionDo.TableName()
	_goadminPermission.ALL = field.NewAsterisk(tableName)
	_goadminPermission.ID = field.NewInt32(tableName, "id")
	_goadminPermission.Name = field.NewString(tableName, "name")
	_goadminPermission.Slug = field.NewString(tableName, "slug")
	_goadminPermission.HTTPMethod = field.NewString(tableName, "http_method")
	_goadminPermission.HTTPPath = field.NewString(tableName, "http_path")
	_goadminPermission.CreatedAt = field.NewTime(tableName, "created_at")
	_goadminPermission.UpdatedAt = field.NewTime(tableName, "updated_at")

	_goadminPermission.fillFieldMap()

	return _goadminPermission
}

type goadminPermission struct {
	goadminPermissionDo goadminPermissionDo

	ALL        field.Asterisk
	ID         field.Int32
	Name       field.String
	Slug       field.String
	HTTPMethod field.String
	HTTPPath   field.String
	CreatedAt  field.Time
	UpdatedAt  field.Time

	fieldMap map[string]field.Expr
}

func (g goadminPermission) Table(newTableName string) *goadminPermission {
	g.goadminPermissionDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g goadminPermission) As(alias string) *goadminPermission {
	g.goadminPermissionDo.DO = *(g.goadminPermissionDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *goadminPermission) updateTableName(table string) *goadminPermission {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt32(table, "id")
	g.Name = field.NewString(table, "name")
	g.Slug = field.NewString(table, "slug")
	g.HTTPMethod = field.NewString(table, "http_method")
	g.HTTPPath = field.NewString(table, "http_path")
	g.CreatedAt = field.NewTime(table, "created_at")
	g.UpdatedAt = field.NewTime(table, "updated_at")

	g.fillFieldMap()

	return g
}

func (g *goadminPermission) WithContext(ctx context.Context) IGoadminPermissionDo {
	return g.goadminPermissionDo.WithContext(ctx)
}

func (g goadminPermission) TableName() string { return g.goadminPermissionDo.TableName() }

func (g goadminPermission) Alias() string { return g.goadminPermissionDo.Alias() }

func (g *goadminPermission) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *goadminPermission) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 7)
	g.fieldMap["id"] = g.ID
	g.fieldMap["name"] = g.Name
	g.fieldMap["slug"] = g.Slug
	g.fieldMap["http_method"] = g.HTTPMethod
	g.fieldMap["http_path"] = g.HTTPPath
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
}

func (g goadminPermission) clone(db *gorm.DB) goadminPermission {
	g.goadminPermissionDo.ReplaceDB(db)
	return g
}

type goadminPermissionDo struct{ gen.DO }

type IGoadminPermissionDo interface {
	gen.SubQuery
	Debug() IGoadminPermissionDo
	WithContext(ctx context.Context) IGoadminPermissionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGoadminPermissionDo
	WriteDB() IGoadminPermissionDo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGoadminPermissionDo
	Not(conds ...gen.Condition) IGoadminPermissionDo
	Or(conds ...gen.Condition) IGoadminPermissionDo
	Select(conds ...field.Expr) IGoadminPermissionDo
	Where(conds ...gen.Condition) IGoadminPermissionDo
	Order(conds ...field.Expr) IGoadminPermissionDo
	Distinct(cols ...field.Expr) IGoadminPermissionDo
	Omit(cols ...field.Expr) IGoadminPermissionDo
	Join(table schema.Tabler, on ...field.Expr) IGoadminPermissionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGoadminPermissionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGoadminPermissionDo
	Group(cols ...field.Expr) IGoadminPermissionDo
	Having(conds ...gen.Condition) IGoadminPermissionDo
	Limit(limit int) IGoadminPermissionDo
	Offset(offset int) IGoadminPermissionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGoadminPermissionDo
	Unscoped() IGoadminPermissionDo
	Create(values ...*model.GoadminPermission) error
	CreateInBatches(values []*model.GoadminPermission, batchSize int) error
	Save(values ...*model.GoadminPermission) error
	First() (*model.GoadminPermission, error)
	Take() (*model.GoadminPermission, error)
	Last() (*model.GoadminPermission, error)
	Find() ([]*model.GoadminPermission, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GoadminPermission, err error)
	FindInBatches(result *[]*model.GoadminPermission, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.GoadminPermission) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGoadminPermissionDo
	Assign(attrs ...field.AssignExpr) IGoadminPermissionDo
	Joins(fields ...field.RelationField) IGoadminPermissionDo
	Preload(fields ...field.RelationField) IGoadminPermissionDo
	FirstOrInit() (*model.GoadminPermission, error)
	FirstOrCreate() (*model.GoadminPermission, error)
	FindByPage(offset int, limit int) (result []*model.GoadminPermission, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGoadminPermissionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetById(id int) (result *model.GoadminPermission, err error)
}

//SELECT * FROM @@table WHERE id=@id
func (g goadminPermissionDo) GetById(id int) (result *model.GoadminPermission, err error) {
	params := make(map[string]interface{}, 0)

	var generateSQL strings.Builder
	params["id"] = id
	generateSQL.WriteString("SELECT * FROM goadmin_permissions WHERE id=@id ")

	var executeSQL *gorm.DB
	if len(params) > 0 {
		executeSQL = g.UnderlyingDB().Raw(generateSQL.String(), params).Take(&result)
	} else {
		executeSQL = g.UnderlyingDB().Raw(generateSQL.String()).Take(&result)
	}
	err = executeSQL.Error
	return
}

func (g goadminPermissionDo) Debug() IGoadminPermissionDo {
	return g.withDO(g.DO.Debug())
}

func (g goadminPermissionDo) WithContext(ctx context.Context) IGoadminPermissionDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g goadminPermissionDo) ReadDB() IGoadminPermissionDo {
	return g.Clauses(dbresolver.Read)
}

func (g goadminPermissionDo) WriteDB() IGoadminPermissionDo {
	return g.Clauses(dbresolver.Write)
}

func (g goadminPermissionDo) Clauses(conds ...clause.Expression) IGoadminPermissionDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g goadminPermissionDo) Returning(value interface{}, columns ...string) IGoadminPermissionDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g goadminPermissionDo) Not(conds ...gen.Condition) IGoadminPermissionDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g goadminPermissionDo) Or(conds ...gen.Condition) IGoadminPermissionDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g goadminPermissionDo) Select(conds ...field.Expr) IGoadminPermissionDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g goadminPermissionDo) Where(conds ...gen.Condition) IGoadminPermissionDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g goadminPermissionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IGoadminPermissionDo {
	return g.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (g goadminPermissionDo) Order(conds ...field.Expr) IGoadminPermissionDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g goadminPermissionDo) Distinct(cols ...field.Expr) IGoadminPermissionDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g goadminPermissionDo) Omit(cols ...field.Expr) IGoadminPermissionDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g goadminPermissionDo) Join(table schema.Tabler, on ...field.Expr) IGoadminPermissionDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g goadminPermissionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGoadminPermissionDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g goadminPermissionDo) RightJoin(table schema.Tabler, on ...field.Expr) IGoadminPermissionDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g goadminPermissionDo) Group(cols ...field.Expr) IGoadminPermissionDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g goadminPermissionDo) Having(conds ...gen.Condition) IGoadminPermissionDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g goadminPermissionDo) Limit(limit int) IGoadminPermissionDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g goadminPermissionDo) Offset(offset int) IGoadminPermissionDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g goadminPermissionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGoadminPermissionDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g goadminPermissionDo) Unscoped() IGoadminPermissionDo {
	return g.withDO(g.DO.Unscoped())
}

func (g goadminPermissionDo) Create(values ...*model.GoadminPermission) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g goadminPermissionDo) CreateInBatches(values []*model.GoadminPermission, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g goadminPermissionDo) Save(values ...*model.GoadminPermission) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g goadminPermissionDo) First() (*model.GoadminPermission, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminPermission), nil
	}
}

func (g goadminPermissionDo) Take() (*model.GoadminPermission, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminPermission), nil
	}
}

func (g goadminPermissionDo) Last() (*model.GoadminPermission, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminPermission), nil
	}
}

func (g goadminPermissionDo) Find() ([]*model.GoadminPermission, error) {
	result, err := g.DO.Find()
	return result.([]*model.GoadminPermission), err
}

func (g goadminPermissionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GoadminPermission, err error) {
	buf := make([]*model.GoadminPermission, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g goadminPermissionDo) FindInBatches(result *[]*model.GoadminPermission, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g goadminPermissionDo) Attrs(attrs ...field.AssignExpr) IGoadminPermissionDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g goadminPermissionDo) Assign(attrs ...field.AssignExpr) IGoadminPermissionDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g goadminPermissionDo) Joins(fields ...field.RelationField) IGoadminPermissionDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g goadminPermissionDo) Preload(fields ...field.RelationField) IGoadminPermissionDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g goadminPermissionDo) FirstOrInit() (*model.GoadminPermission, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminPermission), nil
	}
}

func (g goadminPermissionDo) FirstOrCreate() (*model.GoadminPermission, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoadminPermission), nil
	}
}

func (g goadminPermissionDo) FindByPage(offset int, limit int) (result []*model.GoadminPermission, count int64, err error) {
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

func (g goadminPermissionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g goadminPermissionDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g goadminPermissionDo) Delete(models ...*model.GoadminPermission) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *goadminPermissionDo) withDO(do gen.Dao) *goadminPermissionDo {
	g.DO = *do.(*gen.DO)
	return g
}
