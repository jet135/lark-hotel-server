// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LarkEventOperationLogDao is the data access object for table lark_event_operation_log.
type LarkEventOperationLogDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of current DAO.
	columns LarkEventOperationLogColumns // columns contains all the column names of Table for convenient usage.
}

// LarkEventOperationLogColumns defines and stores column names for table lark_event_operation_log.
type LarkEventOperationLogColumns struct {
	Id        string //
	MsgInfo   string //
	EventType string //
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// larkEventOperationLogColumns holds the columns for table lark_event_operation_log.
var larkEventOperationLogColumns = LarkEventOperationLogColumns{
	Id:        "id",
	MsgInfo:   "msg_info",
	EventType: "event_type",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewLarkEventOperationLogDao creates and returns a new DAO object for table data access.
func NewLarkEventOperationLogDao() *LarkEventOperationLogDao {
	return &LarkEventOperationLogDao{
		group:   "default",
		table:   "lark_event_operation_log",
		columns: larkEventOperationLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LarkEventOperationLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LarkEventOperationLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LarkEventOperationLogDao) Columns() LarkEventOperationLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LarkEventOperationLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LarkEventOperationLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LarkEventOperationLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
