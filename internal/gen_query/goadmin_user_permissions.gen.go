// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen_query

import (
	"context"
	"jim_service/internal/gen_model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newGoadminUserPermissions(db *gorm.DB) goadminUserPermissions {
	_goadminUserPermissions := goadminUserPermissions{}

	_goadminUserPermissions.goadminUserPermissionsDo.UseDB(db)
	_goadminUserPermissions.goadminUserPermissionsDo.UseModel(&gen_model.GoadminUserPermissions{})

	tableName := _goadminUserPermissions.goadminUserPermissionsDo.TableName()
	_goadminUserPermissions.ALL = field.NewField(tableName, "*")
	_goadminUserPermissions.UserID = field.NewInt32(tableName, "user_id")
	_goadminUserPermissions.PermissionID = field.NewInt32(tableName, "permission_id")
	_goadminUserPermissions.CreatedAt = field.NewTime(tableName, "created_at")
	_goadminUserPermissions.UpdatedAt = field.NewTime(tableName, "updated_at")

	_goadminUserPermissions.fillFieldMap()

	return _goadminUserPermissions
}

type goadminUserPermissions struct {
	goadminUserPermissionsDo goadminUserPermissionsDo

	ALL          field.Field
	UserID       field.Int32
	PermissionID field.Int32
	CreatedAt    field.Time
	UpdatedAt    field.Time

	fieldMap map[string]field.Expr
}

func (g goadminUserPermissions) As(alias string) *goadminUserPermissions {
	g.goadminUserPermissionsDo.DO = *(g.goadminUserPermissionsDo.As(alias).(*gen.DO))

	g.ALL = field.NewField(alias, "*")
	g.UserID = field.NewInt32(alias, "user_id")
	g.PermissionID = field.NewInt32(alias, "permission_id")
	g.CreatedAt = field.NewTime(alias, "created_at")
	g.UpdatedAt = field.NewTime(alias, "updated_at")

	g.fillFieldMap()

	return &g
}

func (g *goadminUserPermissions) WithContext(ctx context.Context) *goadminUserPermissionsDo {
	return g.goadminUserPermissionsDo.WithContext(ctx)
}

func (g goadminUserPermissions) TableName() string { return g.goadminUserPermissionsDo.TableName() }

func (g *goadminUserPermissions) GetFieldByName(fieldName string) (field.Expr, bool) {
	field, ok := g.fieldMap[fieldName]
	return field, ok
}

func (g *goadminUserPermissions) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 4)
	g.fieldMap["user_id"] = g.UserID
	g.fieldMap["permission_id"] = g.PermissionID
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
}

func (g goadminUserPermissions) clone(db *gorm.DB) goadminUserPermissions {
	g.goadminUserPermissionsDo.ReplaceDB(db)
	return g
}

type goadminUserPermissionsDo struct{ gen.DO }

func (g goadminUserPermissionsDo) Debug() *goadminUserPermissionsDo {
	return g.withDO(g.DO.Debug())
}

func (g goadminUserPermissionsDo) WithContext(ctx context.Context) *goadminUserPermissionsDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g goadminUserPermissionsDo) Clauses(conds ...clause.Expression) *goadminUserPermissionsDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g goadminUserPermissionsDo) Not(conds ...gen.Condition) *goadminUserPermissionsDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g goadminUserPermissionsDo) Or(conds ...gen.Condition) *goadminUserPermissionsDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g goadminUserPermissionsDo) Select(conds ...field.Expr) *goadminUserPermissionsDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g goadminUserPermissionsDo) Where(conds ...gen.Condition) *goadminUserPermissionsDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g goadminUserPermissionsDo) Order(conds ...field.Expr) *goadminUserPermissionsDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g goadminUserPermissionsDo) Distinct(cols ...field.Expr) *goadminUserPermissionsDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g goadminUserPermissionsDo) Omit(cols ...field.Expr) *goadminUserPermissionsDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g goadminUserPermissionsDo) Join(table schema.Tabler, on ...field.Expr) *goadminUserPermissionsDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g goadminUserPermissionsDo) LeftJoin(table schema.Tabler, on ...field.Expr) *goadminUserPermissionsDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g goadminUserPermissionsDo) RightJoin(table schema.Tabler, on ...field.Expr) *goadminUserPermissionsDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g goadminUserPermissionsDo) Group(cols ...field.Expr) *goadminUserPermissionsDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g goadminUserPermissionsDo) Having(conds ...gen.Condition) *goadminUserPermissionsDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g goadminUserPermissionsDo) Limit(limit int) *goadminUserPermissionsDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g goadminUserPermissionsDo) Offset(offset int) *goadminUserPermissionsDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g goadminUserPermissionsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *goadminUserPermissionsDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g goadminUserPermissionsDo) Unscoped() *goadminUserPermissionsDo {
	return g.withDO(g.DO.Unscoped())
}

func (g goadminUserPermissionsDo) Create(values ...*gen_model.GoadminUserPermissions) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g goadminUserPermissionsDo) CreateInBatches(values []*gen_model.GoadminUserPermissions, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g goadminUserPermissionsDo) Save(values ...*gen_model.GoadminUserPermissions) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g goadminUserPermissionsDo) First() (*gen_model.GoadminUserPermissions, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminUserPermissions), nil
	}
}

func (g goadminUserPermissionsDo) Take() (*gen_model.GoadminUserPermissions, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminUserPermissions), nil
	}
}

func (g goadminUserPermissionsDo) Last() (*gen_model.GoadminUserPermissions, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminUserPermissions), nil
	}
}

func (g goadminUserPermissionsDo) Find() ([]*gen_model.GoadminUserPermissions, error) {
	result, err := g.DO.Find()
	return result.([]*gen_model.GoadminUserPermissions), err
}

func (g goadminUserPermissionsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gen_model.GoadminUserPermissions, err error) {
	buf := make([]*gen_model.GoadminUserPermissions, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g goadminUserPermissionsDo) FindInBatches(result *[]*gen_model.GoadminUserPermissions, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g goadminUserPermissionsDo) Attrs(attrs ...field.AssignExpr) *goadminUserPermissionsDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g goadminUserPermissionsDo) Assign(attrs ...field.AssignExpr) *goadminUserPermissionsDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g goadminUserPermissionsDo) Joins(field field.RelationField) *goadminUserPermissionsDo {
	return g.withDO(g.DO.Joins(field))
}

func (g goadminUserPermissionsDo) Preload(field field.RelationField) *goadminUserPermissionsDo {
	return g.withDO(g.DO.Preload(field))
}

func (g goadminUserPermissionsDo) FirstOrInit() (*gen_model.GoadminUserPermissions, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminUserPermissions), nil
	}
}

func (g goadminUserPermissionsDo) FirstOrCreate() (*gen_model.GoadminUserPermissions, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*gen_model.GoadminUserPermissions), nil
	}
}

func (g goadminUserPermissionsDo) FindByPage(offset int, limit int) (result []*gen_model.GoadminUserPermissions, count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	result, err = g.Offset(offset).Limit(limit).Find()
	return
}

func (g goadminUserPermissionsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g *goadminUserPermissionsDo) withDO(do gen.Dao) *goadminUserPermissionsDo {
	g.DO = *do.(*gen.DO)
	return g
}