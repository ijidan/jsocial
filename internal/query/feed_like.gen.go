// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

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

func (f *feedLike) WithContext(ctx context.Context) *feedLikeDo { return f.feedLikeDo.WithContext(ctx) }

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

func (f feedLikeDo) Debug() *feedLikeDo {
	return f.withDO(f.DO.Debug())
}

func (f feedLikeDo) WithContext(ctx context.Context) *feedLikeDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f feedLikeDo) ReadDB() *feedLikeDo {
	return f.Clauses(dbresolver.Read)
}

func (f feedLikeDo) WriteDB() *feedLikeDo {
	return f.Clauses(dbresolver.Write)
}

func (f feedLikeDo) Clauses(conds ...clause.Expression) *feedLikeDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f feedLikeDo) Returning(value interface{}, columns ...string) *feedLikeDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f feedLikeDo) Not(conds ...gen.Condition) *feedLikeDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f feedLikeDo) Or(conds ...gen.Condition) *feedLikeDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f feedLikeDo) Select(conds ...field.Expr) *feedLikeDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f feedLikeDo) Where(conds ...gen.Condition) *feedLikeDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f feedLikeDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *feedLikeDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f feedLikeDo) Order(conds ...field.Expr) *feedLikeDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f feedLikeDo) Distinct(cols ...field.Expr) *feedLikeDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f feedLikeDo) Omit(cols ...field.Expr) *feedLikeDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f feedLikeDo) Join(table schema.Tabler, on ...field.Expr) *feedLikeDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f feedLikeDo) LeftJoin(table schema.Tabler, on ...field.Expr) *feedLikeDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f feedLikeDo) RightJoin(table schema.Tabler, on ...field.Expr) *feedLikeDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f feedLikeDo) Group(cols ...field.Expr) *feedLikeDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f feedLikeDo) Having(conds ...gen.Condition) *feedLikeDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f feedLikeDo) Limit(limit int) *feedLikeDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f feedLikeDo) Offset(offset int) *feedLikeDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f feedLikeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *feedLikeDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f feedLikeDo) Unscoped() *feedLikeDo {
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

func (f feedLikeDo) Attrs(attrs ...field.AssignExpr) *feedLikeDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f feedLikeDo) Assign(attrs ...field.AssignExpr) *feedLikeDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f feedLikeDo) Joins(fields ...field.RelationField) *feedLikeDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f feedLikeDo) Preload(fields ...field.RelationField) *feedLikeDo {
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