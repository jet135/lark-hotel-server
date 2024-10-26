// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReportLogsDao is the data access object for table report_logs.
type ReportLogsDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns ReportLogsColumns // columns contains all the column names of Table for convenient usage.
}

// ReportLogsColumns defines and stores column names for table report_logs.
type ReportLogsColumns struct {
	Id          string //
	OriginalUrl string //
	Method      string //
	RequestBody string //
	CreatedAt   string //
	UpdatedAt   string //
	DeletedAt   string //
}

// reportLogsColumns holds the columns for table report_logs.
var reportLogsColumns = ReportLogsColumns{
	Id:          "id",
	OriginalUrl: "original_url",
	Method:      "method",
	RequestBody: "request_body",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewReportLogsDao creates and returns a new DAO object for table data access.
func NewReportLogsDao() *ReportLogsDao {
	return &ReportLogsDao{
		group:   "default",
		table:   "report_logs",
		columns: reportLogsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ReportLogsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ReportLogsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ReportLogsDao) Columns() ReportLogsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ReportLogsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ReportLogsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ReportLogsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
