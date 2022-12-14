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

func newFeedLike(db *gorm.DB) feedLike {
	_feedLike := feedLike{}

	_feedLike.feedLikeDo.UseDB(db)
	_feedLike.feedLikeDo.UseModel(&model.FeedLike{})

	tableName := _feedLike.feedLikeDo.TableName()
	_feedLike.ALL = field.NewAsterisk(tableName)
	_feedLike.ID = field.NewInt64(tableName, "id")
	_feedLike.FeedID = field.NewInt64(tableName, "feed_id")
	_feedLike.UserID = field.NewInt64(tableName, "user_id")
	_feedLike.CreatedAt = field.NewTime(tableName, "created_at")
	_feedLike.UpdatedAt = field.NewTime(tableName, "updated_at")
	_feedLike.DeletedAt = field.NewField(tableName, "deleted_at")

	_feedLike.fillFieldMap()

	return _feedLike
}

type feedLike struct {
	feedLikeDo feedLikeDo

	ALL       field.Asterisk
	ID        field.Int64 // ID
	FeedID    field.Int64 // 动态ID
	UserID    field.Int64 // 用户ID
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (f feedLike) Table(newTableName string) *feedLike {
	f.feedLikeDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f feedLike) As(alias string) *feedLike {
	f.feedLikeDo.DO = *(f.feedLikeDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *feedLike) updateTableName(table string) *feedLike {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.FeedID = field.NewInt64(table, "feed_id")
	f.UserID = field.NewInt64(table, "user_id")
	f.CreatedAt = field.NewTime(table, "created_at")
	f.UpdatedAt = field.NewTime(table, "updated_at")
	f.DeletedAt = field.NewField(table, "deleted_at")

	f.fillFieldMap()

	return f
}

func (f *feedLike) WithContext(ctx context.Context) IFeedLikeDo { return f.feedLikeDo.WithContext(ctx) }

func (f feedLike) TableName() string { return f.feedLikeDo.TableName() }

func (f feedLike) Alias() string { return f.feedLikeDo.Alias() }

func (f *feedLike) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *feedLike) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 6)
	f.fieldMap["id"] = f.ID
	f.fieldMap["feed_id"] = f.FeedID
	f.fieldMap["user_id"] = f.UserID
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["deleted_at"] = f.DeletedAt
}

func (f feedLike) clone(db *gorm.DB) feedLike {
	f.feedLikeDo.ReplaceDB(db)
	return f
}

type feedLikeDo struct{ gen.DO }

type IFeedLikeDo interface {
	gen.SubQuery
	Debug() IFeedLikeDo
	WithContext(ctx context.Context) IFeedLikeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFeedLikeDo
	WriteDB() IFeedLikeDo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFeedLikeDo
	Not(conds ...gen.Condition) IFeedLikeDo
	Or(conds ...gen.Condition) IFeedLikeDo
	Select(conds ...field.Expr) IFeedLikeDo
	Where(conds ...gen.Condition) IFeedLikeDo
	Order(conds ...field.Expr) IFeedLikeDo
	Distinct(cols ...field.Expr) IFeedLikeDo
	Omit(cols ...field.Expr) IFeedLikeDo
	Join(table schema.Tabler, on ...field.Expr) IFeedLikeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFeedLikeDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFeedLikeDo
	Group(cols ...field.Expr) IFeedLikeDo
	Having(conds ...gen.Condition) IFeedLikeDo
	Limit(limit int) IFeedLikeDo
	Offset(offset int) IFeedLikeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFeedLikeDo
	Unscoped() IFeedLikeDo
	Create(values ...*model.FeedLike) error
	CreateInBatches(values []*model.FeedLike, batchSize int) error
	Save(values ...*model.FeedLike) error
	First() (*model.FeedLike, error)
	Take() (*model.FeedLike, error)
	Last() (*model.FeedLike, error)
	Find() ([]*model.FeedLike, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FeedLike, err error)
	FindInBatches(result *[]*model.FeedLike, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FeedLike) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFeedLikeDo
	Assign(attrs ...field.AssignExpr) IFeedLikeDo
	Joins(fields ...field.RelationField) IFeedLikeDo
	Preload(fields ...field.RelationField) IFeedLikeDo
	FirstOrInit() (*model.FeedLike, error)
	FirstOrCreate() (*model.FeedLike, error)
	FindByPage(offset int, limit int) (result []*model.FeedLike, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFeedLikeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetById(id int) (result *model.FeedLike, err error)
}

//SELECT * FROM @@table WHERE id=@id
func (f feedLikeDo) GetById(id int) (result *model.FeedLike, err error) {
	params := make(map[string]interface{}, 0)

	var generateSQL strings.Builder
	params["id"] = id
	generateSQL.WriteString("SELECT * FROM feed_like WHERE id=@id ")

	var executeSQL *gorm.DB
	if len(params) > 0 {
		executeSQL = f.UnderlyingDB().Raw(generateSQL.String(), params).Take(&result)
	} else {
		executeSQL = f.UnderlyingDB().Raw(generateSQL.String()).Take(&result)
	}
	err = executeSQL.Error
	return
}

func (f feedLikeDo) Debug() IFeedLikeDo {
	return f.withDO(f.DO.Debug())
}

func (f feedLikeDo) WithContext(ctx context.Context) IFeedLikeDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f feedLikeDo) ReadDB() IFeedLikeDo {
	return f.Clauses(dbresolver.Read)
}

func (f feedLikeDo) WriteDB() IFeedLikeDo {
	return f.Clauses(dbresolver.Write)
}

func (f feedLikeDo) Clauses(conds ...clause.Expression) IFeedLikeDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f feedLikeDo) Returning(value interface{}, columns ...string) IFeedLikeDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f feedLikeDo) Not(conds ...gen.Condition) IFeedLikeDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f feedLikeDo) Or(conds ...gen.Condition) IFeedLikeDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f feedLikeDo) Select(conds ...field.Expr) IFeedLikeDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f feedLikeDo) Where(conds ...gen.Condition) IFeedLikeDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f feedLikeDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IFeedLikeDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f feedLikeDo) Order(conds ...field.Expr) IFeedLikeDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f feedLikeDo) Distinct(cols ...field.Expr) IFeedLikeDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f feedLikeDo) Omit(cols ...field.Expr) IFeedLikeDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f feedLikeDo) Join(table schema.Tabler, on ...field.Expr) IFeedLikeDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f feedLikeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFeedLikeDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f feedLikeDo) RightJoin(table schema.Tabler, on ...field.Expr) IFeedLikeDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f feedLikeDo) Group(cols ...field.Expr) IFeedLikeDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f feedLikeDo) Having(conds ...gen.Condition) IFeedLikeDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f feedLikeDo) Limit(limit int) IFeedLikeDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f feedLikeDo) Offset(offset int) IFeedLikeDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f feedLikeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFeedLikeDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f feedLikeDo) Unscoped() IFeedLikeDo {
	return f.withDO(f.DO.Unscoped())
}

func (f feedLikeDo) Create(values ...*model.FeedLike) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f feedLikeDo) CreateInBatches(values []*model.FeedLike, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f feedLikeDo) Save(values ...*model.FeedLike) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f feedLikeDo) First() (*model.FeedLike, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FeedLike), nil
	}
}

func (f feedLikeDo) Take() (*model.FeedLike, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FeedLike), nil
	}
}

func (f feedLikeDo) Last() (*model.FeedLike, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FeedLike), nil
	}
}

func (f feedLikeDo) Find() ([]*model.FeedLike, error) {
	result, err := f.DO.Find()
	return result.([]*model.FeedLike), err
}

func (f feedLikeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FeedLike, err error) {
	buf := make([]*model.FeedLike, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f feedLikeDo) FindInBatches(result *[]*model.FeedLike, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f feedLikeDo) Attrs(attrs ...field.AssignExpr) IFeedLikeDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f feedLikeDo) Assign(attrs ...field.AssignExpr) IFeedLikeDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f feedLikeDo) Joins(fields ...field.RelationField) IFeedLikeDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f feedLikeDo) Preload(fields ...field.RelationField) IFeedLikeDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f feedLikeDo) FirstOrInit() (*model.FeedLike, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FeedLike), nil
	}
}

func (f feedLikeDo) FirstOrCreate() (*model.FeedLike, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FeedLike), nil
	}
}

func (f feedLikeDo) FindByPage(offset int, limit int) (result []*model.FeedLike, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f feedLikeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f feedLikeDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f feedLikeDo) Delete(models ...*model.FeedLike) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *feedLikeDo) withDO(do gen.Dao) *feedLikeDo {
	f.DO = *do.(*gen.DO)
	return f
}
