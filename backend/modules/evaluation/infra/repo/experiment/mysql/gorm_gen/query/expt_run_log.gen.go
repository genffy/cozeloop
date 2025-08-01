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

	"github.com/coze-dev/coze-loop/backend/modules/evaluation/infra/repo/experiment/mysql/gorm_gen/model"
)

func newExptRunLog(db *gorm.DB, opts ...gen.DOOption) exptRunLog {
	_exptRunLog := exptRunLog{}

	_exptRunLog.exptRunLogDo.UseDB(db, opts...)
	_exptRunLog.exptRunLogDo.UseModel(&model.ExptRunLog{})

	tableName := _exptRunLog.exptRunLogDo.TableName()
	_exptRunLog.ALL = field.NewAsterisk(tableName)
	_exptRunLog.ID = field.NewInt64(tableName, "id")
	_exptRunLog.SpaceID = field.NewInt64(tableName, "space_id")
	_exptRunLog.CreatedBy = field.NewString(tableName, "created_by")
	_exptRunLog.ExptID = field.NewInt64(tableName, "expt_id")
	_exptRunLog.ExptRunID = field.NewInt64(tableName, "expt_run_id")
	_exptRunLog.ItemIds = field.NewBytes(tableName, "item_ids")
	_exptRunLog.Mode = field.NewInt32(tableName, "mode")
	_exptRunLog.Status = field.NewInt64(tableName, "status")
	_exptRunLog.PendingCnt = field.NewInt32(tableName, "pending_cnt")
	_exptRunLog.SuccessCnt = field.NewInt32(tableName, "success_cnt")
	_exptRunLog.FailCnt = field.NewInt32(tableName, "fail_cnt")
	_exptRunLog.CreditCost = field.NewFloat64(tableName, "credit_cost")
	_exptRunLog.TokenCost = field.NewInt64(tableName, "token_cost")
	_exptRunLog.CreatedAt = field.NewTime(tableName, "created_at")
	_exptRunLog.UpdatedAt = field.NewTime(tableName, "updated_at")
	_exptRunLog.DeletedAt = field.NewField(tableName, "deleted_at")
	_exptRunLog.StatusMessage = field.NewBytes(tableName, "status_message")
	_exptRunLog.ProcessingCnt = field.NewInt32(tableName, "processing_cnt")
	_exptRunLog.TerminatedCnt = field.NewInt32(tableName, "terminated_cnt")

	_exptRunLog.fillFieldMap()

	return _exptRunLog
}

// exptRunLog expt_run_log
type exptRunLog struct {
	exptRunLogDo exptRunLogDo

	ALL           field.Asterisk
	ID            field.Int64   // id
	SpaceID       field.Int64   // 空间 id
	CreatedBy     field.String  // 创建者 id
	ExptID        field.Int64   // 实验 id
	ExptRunID     field.Int64   // 运行 id
	ItemIds       field.Bytes   // 组 ids
	Mode          field.Int32   // 模式
	Status        field.Int64   // 状态
	PendingCnt    field.Int32   // item 未执行数量
	SuccessCnt    field.Int32   // item 成功数量
	FailCnt       field.Int32   // item 失败数量
	CreditCost    field.Float64 // credit 消耗
	TokenCost     field.Int64   // token 消耗
	CreatedAt     field.Time    // 创建时间
	UpdatedAt     field.Time    // 更新时间
	DeletedAt     field.Field   // 删除时间
	StatusMessage field.Bytes   // 提示信息
	ProcessingCnt field.Int32   // processing_cnt
	TerminatedCnt field.Int32   // terminated_cnt

	fieldMap map[string]field.Expr
}

func (e exptRunLog) Table(newTableName string) *exptRunLog {
	e.exptRunLogDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e exptRunLog) As(alias string) *exptRunLog {
	e.exptRunLogDo.DO = *(e.exptRunLogDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *exptRunLog) updateTableName(table string) *exptRunLog {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt64(table, "id")
	e.SpaceID = field.NewInt64(table, "space_id")
	e.CreatedBy = field.NewString(table, "created_by")
	e.ExptID = field.NewInt64(table, "expt_id")
	e.ExptRunID = field.NewInt64(table, "expt_run_id")
	e.ItemIds = field.NewBytes(table, "item_ids")
	e.Mode = field.NewInt32(table, "mode")
	e.Status = field.NewInt64(table, "status")
	e.PendingCnt = field.NewInt32(table, "pending_cnt")
	e.SuccessCnt = field.NewInt32(table, "success_cnt")
	e.FailCnt = field.NewInt32(table, "fail_cnt")
	e.CreditCost = field.NewFloat64(table, "credit_cost")
	e.TokenCost = field.NewInt64(table, "token_cost")
	e.CreatedAt = field.NewTime(table, "created_at")
	e.UpdatedAt = field.NewTime(table, "updated_at")
	e.DeletedAt = field.NewField(table, "deleted_at")
	e.StatusMessage = field.NewBytes(table, "status_message")
	e.ProcessingCnt = field.NewInt32(table, "processing_cnt")
	e.TerminatedCnt = field.NewInt32(table, "terminated_cnt")

	e.fillFieldMap()

	return e
}

func (e *exptRunLog) WithContext(ctx context.Context) *exptRunLogDo {
	return e.exptRunLogDo.WithContext(ctx)
}

func (e exptRunLog) TableName() string { return e.exptRunLogDo.TableName() }

func (e exptRunLog) Alias() string { return e.exptRunLogDo.Alias() }

func (e exptRunLog) Columns(cols ...field.Expr) gen.Columns { return e.exptRunLogDo.Columns(cols...) }

func (e *exptRunLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *exptRunLog) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 19)
	e.fieldMap["id"] = e.ID
	e.fieldMap["space_id"] = e.SpaceID
	e.fieldMap["created_by"] = e.CreatedBy
	e.fieldMap["expt_id"] = e.ExptID
	e.fieldMap["expt_run_id"] = e.ExptRunID
	e.fieldMap["item_ids"] = e.ItemIds
	e.fieldMap["mode"] = e.Mode
	e.fieldMap["status"] = e.Status
	e.fieldMap["pending_cnt"] = e.PendingCnt
	e.fieldMap["success_cnt"] = e.SuccessCnt
	e.fieldMap["fail_cnt"] = e.FailCnt
	e.fieldMap["credit_cost"] = e.CreditCost
	e.fieldMap["token_cost"] = e.TokenCost
	e.fieldMap["created_at"] = e.CreatedAt
	e.fieldMap["updated_at"] = e.UpdatedAt
	e.fieldMap["deleted_at"] = e.DeletedAt
	e.fieldMap["status_message"] = e.StatusMessage
	e.fieldMap["processing_cnt"] = e.ProcessingCnt
	e.fieldMap["terminated_cnt"] = e.TerminatedCnt
}

func (e exptRunLog) clone(db *gorm.DB) exptRunLog {
	e.exptRunLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e exptRunLog) replaceDB(db *gorm.DB) exptRunLog {
	e.exptRunLogDo.ReplaceDB(db)
	return e
}

type exptRunLogDo struct{ gen.DO }

func (e exptRunLogDo) Debug() *exptRunLogDo {
	return e.withDO(e.DO.Debug())
}

func (e exptRunLogDo) WithContext(ctx context.Context) *exptRunLogDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e exptRunLogDo) ReadDB() *exptRunLogDo {
	return e.Clauses(dbresolver.Read)
}

func (e exptRunLogDo) WriteDB() *exptRunLogDo {
	return e.Clauses(dbresolver.Write)
}

func (e exptRunLogDo) Session(config *gorm.Session) *exptRunLogDo {
	return e.withDO(e.DO.Session(config))
}

func (e exptRunLogDo) Clauses(conds ...clause.Expression) *exptRunLogDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e exptRunLogDo) Returning(value interface{}, columns ...string) *exptRunLogDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e exptRunLogDo) Not(conds ...gen.Condition) *exptRunLogDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e exptRunLogDo) Or(conds ...gen.Condition) *exptRunLogDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e exptRunLogDo) Select(conds ...field.Expr) *exptRunLogDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e exptRunLogDo) Where(conds ...gen.Condition) *exptRunLogDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e exptRunLogDo) Order(conds ...field.Expr) *exptRunLogDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e exptRunLogDo) Distinct(cols ...field.Expr) *exptRunLogDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e exptRunLogDo) Omit(cols ...field.Expr) *exptRunLogDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e exptRunLogDo) Join(table schema.Tabler, on ...field.Expr) *exptRunLogDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e exptRunLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) *exptRunLogDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e exptRunLogDo) RightJoin(table schema.Tabler, on ...field.Expr) *exptRunLogDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e exptRunLogDo) Group(cols ...field.Expr) *exptRunLogDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e exptRunLogDo) Having(conds ...gen.Condition) *exptRunLogDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e exptRunLogDo) Limit(limit int) *exptRunLogDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e exptRunLogDo) Offset(offset int) *exptRunLogDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e exptRunLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *exptRunLogDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e exptRunLogDo) Unscoped() *exptRunLogDo {
	return e.withDO(e.DO.Unscoped())
}

func (e exptRunLogDo) Create(values ...*model.ExptRunLog) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e exptRunLogDo) CreateInBatches(values []*model.ExptRunLog, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e exptRunLogDo) Save(values ...*model.ExptRunLog) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e exptRunLogDo) First() (*model.ExptRunLog, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExptRunLog), nil
	}
}

func (e exptRunLogDo) Take() (*model.ExptRunLog, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExptRunLog), nil
	}
}

func (e exptRunLogDo) Last() (*model.ExptRunLog, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExptRunLog), nil
	}
}

func (e exptRunLogDo) Find() ([]*model.ExptRunLog, error) {
	result, err := e.DO.Find()
	return result.([]*model.ExptRunLog), err
}

func (e exptRunLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ExptRunLog, err error) {
	buf := make([]*model.ExptRunLog, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e exptRunLogDo) FindInBatches(result *[]*model.ExptRunLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e exptRunLogDo) Attrs(attrs ...field.AssignExpr) *exptRunLogDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e exptRunLogDo) Assign(attrs ...field.AssignExpr) *exptRunLogDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e exptRunLogDo) Joins(fields ...field.RelationField) *exptRunLogDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e exptRunLogDo) Preload(fields ...field.RelationField) *exptRunLogDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e exptRunLogDo) FirstOrInit() (*model.ExptRunLog, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExptRunLog), nil
	}
}

func (e exptRunLogDo) FirstOrCreate() (*model.ExptRunLog, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExptRunLog), nil
	}
}

func (e exptRunLogDo) FindByPage(offset int, limit int) (result []*model.ExptRunLog, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e exptRunLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e exptRunLogDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e exptRunLogDo) Delete(models ...*model.ExptRunLog) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *exptRunLogDo) withDO(do gen.Dao) *exptRunLogDo {
	e.DO = *do.(*gen.DO)
	return e
}
