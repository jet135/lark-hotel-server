// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReportCustomerDao is the data access object for table report_customer.
type ReportCustomerDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns ReportCustomerColumns // columns contains all the column names of Table for convenient usage.
}

// ReportCustomerColumns defines and stores column names for table report_customer.
type ReportCustomerColumns struct {
	Id         string //
	Name       string //
	BirthDate  string //
	Czdz       string //
	IdCode     string //
	IdType     string //
	Phone      string //
	XzqhTitle  string //
	RoomNumber string //
	CreatedAt  string //
	UpdatedAt  string //
	DeletedAt  string //
}

// reportCustomerColumns holds the columns for table report_customer.
var reportCustomerColumns = ReportCustomerColumns{
	Id:         "id",
	Name:       "name",
	BirthDate:  "birth_date",
	Czdz:       "czdz",
	IdCode:     "id_code",
	IdType:     "id_type",
	Phone:      "phone",
	XzqhTitle:  "xzqhTitle",
	RoomNumber: "room_number",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// NewReportCustomerDao creates and returns a new DAO object for table data access.
func NewReportCustomerDao() *ReportCustomerDao {
	return &ReportCustomerDao{
		group:   "default",
		table:   "report_customer",
		columns: reportCustomerColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ReportCustomerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ReportCustomerDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ReportCustomerDao) Columns() ReportCustomerColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ReportCustomerDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ReportCustomerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ReportCustomerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
