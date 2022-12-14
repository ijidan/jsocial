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

func newFeedVideo(db *gorm.DB) feedVideo {
	_feedVideo := feedVideo{}

	_feedVideo.feedVideoDo.UseDB(db)
	_feedVideo.feedVideoDo.UseModel(&model.FeedVideo{})

	tableName := _feedVideo.feedVideoDo.TableName()
	_feedVideo.ALL = field.NewAsterisk(tableName)
	_feedVideo.ID = field.NewInt32(tableName, "id")
	_feedVideo.FeedID = field.NewInt32(tableName, "feed_id")
	_feedVideo.VideoURL = field.NewString(tableName, "video_url")
	_feedVideo.CoverURL = field.NewString(tableName, "cover_url")
	_feedVideo.Height = field.NewInt32(tableName, "height")
	_feedVideo.Width = field.NewInt32(tableName, "width")
	_feedVideo.Duration = field.NewFloat32(tableName, "duration")
	_feedVideo.CreatedAt = field.NewTime(tableName, "created_at")
	_feedVideo.UpdatedAt = field.NewTime(tableName, "updated_at")
	_feedVideo.DeletedAt = field.NewField(tableName, "deleted_at")

	_feedVideo.fillFieldMap()

	return _feedVideo
}

type feedVideo struct {
	feedVideoDo feedVideoDo

	ALL       field.Asterisk
	ID        field.Int32
	FeedID    field.Int32   // 所属动态id
	VideoURL  field.String  // 视频URL
	CoverURL  field.String  // 视频封面URL
	Height    field.Int32   // 视频高度
	Width     field.Int32   // 视频宽度
	Duration  field.Float32 // 视频时长
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (f feedVideo) Table(newTableName string) *feedVideo {
	f.feedVideoDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f feedVideo) As(alias string) *feedVideo {
	f.feedVideoDo.DO = *(f.feedVideoDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *feedVideo) updateTableName(table string) *feedVideo {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt32(table, "id")
	f.FeedID = field.NewInt32(table, "feed_id")
	f.VideoURL = field.NewString(table, "video_url")
	f.CoverURL = field.NewString(table, "cover_url")
	f.Height = field.NewInt32(table, "height")
	f.Width = field.NewInt32(table, "width")
	f.Duration = field.NewFloat32(table, "duration")
	f.CreatedAt = field.NewTime(table, "created_at")
	f.UpdatedAt = field.NewTime(table, "updated_at")
	f.DeletedAt = field.NewField(table, "deleted_at")

	f.fillFieldMap()

	return f
}

func (f *feedVideo) WithContext(ctx context.Context) IFeedVideoDo {
	return f.feedVideoDo.WithContext(ctx)
}

func (f feedVideo) TableName() string { return f.feedVideoDo.TableName() }

func (f feedVideo) Alias() string { return f.feedVideoDo.Alias() }

func (f *feedVideo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *feedVideo) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 10)
	f.fieldMap["id"] = f.ID
	f.fieldMap["feed_id"] = f.FeedID
	f.fieldMap["video_url"] = f.VideoURL
	f.fieldMap["cover_url"] = f.CoverURL
	f.fieldMap["height"] = f.Height
	f.fieldMap["width"] = f.Width
	f.fieldMap["duration"] = f.Duration
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["deleted_at"] = f.DeletedAt
}

func (f feedVideo) clone(db *gorm.DB) feedVideo {
	f.feedVideoDo.ReplaceDB(db)
	return f
}

type feedVideoDo struct{ gen.DO }

type IFeedVideoDo interface {
	gen.SubQuery
	Debug() IFeedVideoDo
	WithContext(ctx context.Context) IFeedVideoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFeedVideoDo
	WriteDB() IFeedVideoDo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFeedVideoDo
	Not(conds ...gen.Condition) IFeedVideoDo
	Or(conds ...gen.Condition) IFeedVideoDo
	Select(conds ...field.Expr) IFeedVideoDo
	Where(conds ...gen.Condition) IFeedVideoDo
	Order(conds ...field.Expr) IFeedVideoDo
	Distinct(cols ...field.Expr) IFeedVideoDo
	Omit(cols ...field.Expr) IFeedVideoDo
	Join(table schema.Tabler, on ...field.Expr) IFeedVideoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFeedVideoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFeedVideoDo
	Group(cols ...field.Expr) IFeedVideoDo
	Having(conds ...gen.Condition) IFeedVideoDo
	Limit(limit int) IFeedVideoDo
	Offset(offset int) IFeedVideoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFeedVideoDo
	Unscoped() IFeedVideoDo
	Create(values ...*model.FeedVideo) error
	CreateInBatches(values []*model.FeedVideo, batchSize int) error
	Save(values ...*model.FeedVideo) error
	First() (*model.FeedVideo, error)
	Take() (*model.FeedVideo, error)
	Last() (*model.FeedVideo, error)
	Find() ([]*model.FeedVideo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FeedVideo, err error)
	FindInBatches(result *[]*model.FeedVideo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FeedVideo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFeedVideoDo
	Assign(attrs ...field.AssignExpr) IFeedVideoDo
	Joins(fields ...field.RelationField) IFeedVideoDo
	Preload(fields ...field.RelationField) IFeedVideoDo
	FirstOrInit() (*model.FeedVideo, error)
	FirstOrCreate() (*model.FeedVideo, error)
	FindByPage(offset int, limit int) (result []*model.FeedVideo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFeedVideoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetById(id int) (result *model.FeedVideo, err error)
}

//SELECT * FROM @@table WHERE id=@id
func (f feedVideoDo) GetById(id int) (result *model.FeedVideo, err error) {
	params := make(map[string]interface{}, 0)

	var generateSQL strings.Builder
	params["id"] = id
	generateSQL.WriteString("SELECT * FROM feed_video WHERE id=@id ")

	var executeSQL *gorm.DB
	if len(params) > 0 {
		executeSQL = f.UnderlyingDB().Raw(generateSQL.String(), params).Take(&result)
	} else {
		executeSQL = f.UnderlyingDB().Raw(generateSQL.String()).Take(&result)
	}
	err = executeSQL.Error
	return
}

func (f feedVideoDo) Debug() IFeedVideoDo {
	return f.withDO(f.DO.Debug())
}

func (f feedVideoDo) WithContext(ctx context.Context) IFeedVideoDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f feedVideoDo) ReadDB() IFeedVideoDo {
	return f.Clauses(dbresolver.Read)
}

func (f feedVideoDo) WriteDB() IFeedVideoDo {
	return f.Clauses(dbresolver.Write)
}

func (f feedVideoDo) Clauses(conds ...clause.Expression) IFeedVideoDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f feedVideoDo) Returning(value interface{}, columns ...string) IFeedVideoDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f feedVideoDo) Not(conds ...gen.Condition) IFeedVideoDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f feedVideoDo) Or(conds ...gen.Condition) IFeedVideoDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f feedVideoDo) Select(conds ...field.Expr) IFeedVideoDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f feedVideoDo) Where(conds ...gen.Condition) IFeedVideoDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f feedVideoDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IFeedVideoDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f feedVideoDo) Order(conds ...field.Expr) IFeedVideoDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f feedVideoDo) Distinct(cols ...field.Expr) IFeedVideoDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f feedVideoDo) Omit(cols ...field.Expr) IFeedVideoDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f feedVideoDo) Join(table schema.Tabler, on ...field.Expr) IFeedVideoDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f feedVideoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFeedVideoDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f feedVideoDo) RightJoin(table schema.Tabler, on ...field.Expr) IFeedVideoDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f feedVideoDo) Group(cols ...field.Expr) IFeedVideoDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f feedVideoDo) Having(conds ...gen.Condition) IFeedVideoDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f feedVideoDo) Limit(limit int) IFeedVideoDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f feedVideoDo) Offset(offset int) IFeedVideoDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f feedVideoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFeedVideoDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f feedVideoDo) Unscoped() IFeedVideoDo {
	return f.withDO(f.DO.Unscoped())
}

func (f feedVideoDo) Create(values ...*model.FeedVideo) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f feedVideoDo) CreateInBatches(values []*model.FeedVideo, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f feedVideoDo) Save(values ...*model.FeedVideo) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f feedVideoDo) First() (*model.FeedVideo, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FeedVideo), nil
	}
}

func (f feedVideoDo) Take() (*model.FeedVideo, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FeedVideo), nil
	}
}

func (f feedVideoDo) Last() (*model.FeedVideo, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FeedVideo), nil
	}
}

func (f feedVideoDo) Find() ([]*model.FeedVideo, error) {
	result, err := f.DO.Find()
	return result.([]*model.FeedVideo), err
}

func (f feedVideoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FeedVideo, err error) {
	buf := make([]*model.FeedVideo, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f feedVideoDo) FindInBatches(result *[]*model.FeedVideo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f feedVideoDo) Attrs(attrs ...field.AssignExpr) IFeedVideoDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f feedVideoDo) Assign(attrs ...field.AssignExpr) IFeedVideoDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f feedVideoDo) Joins(fields ...field.RelationField) IFeedVideoDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f feedVideoDo) Preload(fields ...field.RelationField) IFeedVideoDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f feedVideoDo) FirstOrInit() (*model.FeedVideo, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FeedVideo), nil
	}
}

func (f feedVideoDo) FirstOrCreate() (*model.FeedVideo, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FeedVideo), nil
	}
}

func (f feedVideoDo) FindByPage(offset int, limit int) (result []*model.FeedVideo, count int64, err error) {
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

func (f feedVideoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f feedVideoDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f feedVideoDo) Delete(models ...*model.FeedVideo) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *feedVideoDo) withDO(do gen.Dao) *feedVideoDo {
	f.DO = *do.(*gen.DO)
	return f
}
