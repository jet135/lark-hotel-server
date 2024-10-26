// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LarkAppInfoDao is the data access object for table lark_app_info.
type LarkAppInfoDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns LarkAppInfoColumns // columns contains all the column names of Table for convenient usage.
}

// LarkAppInfoColumns defines and stores column names for table lark_app_info.
type LarkAppInfoColumns struct {
	Id          string //
	AppName     string //
	AppToken    string //
	FolderToken string //
	Url         string //
	BelongDate  string //
	TableId     string //
	DataSync    string //
	CreatedAt   string //
	UpdatedAt   string //
	DeletedAt   string //
}

// larkAppInfoColumns holds the columns for table lark_app_info.
var larkAppInfoColumns = LarkAppInfoColumns{
	Id:          "id",
	AppName:     "app_name",
	AppToken:    "app_token",
	FolderToken: "folder_token",
	Url:         "url",
	BelongDate:  "belong_date",
	TableId:     "table_id",
	DataSync:    "data_sync",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewLarkAppInfoDao creates and returns a new DAO object for table data access.
func NewLarkAppInfoDao() *LarkAppInfoDao {
	return &LarkAppInfoDao{
		group:   "default",
		table:   "lark_app_info",
		columns: larkAppInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LarkAppInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LarkAppInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LarkAppInfoDao) Columns() LarkAppInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LarkAppInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LarkAppInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LarkAppInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
