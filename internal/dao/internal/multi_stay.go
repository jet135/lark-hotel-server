// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MultiStayDao is the data access object for table multi_stay.
type MultiStayDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns MultiStayColumns // columns contains all the column names of Table for convenient usage.
}

// MultiStayColumns defines and stores column names for table multi_stay.
type MultiStayColumns struct {
	Id                   string //
	CustomerName         string //
	RoomNumber           string //
	PayTime              string //
	CheckinTime          string //
	NumberOfNights       string //
	Done                 string //
	BillId               string //
	ExpectedCheckoutTime string //
	CreatedAt            string //
	UpdatedAt            string //
	DeletedAt            string //
}

// multiStayColumns holds the columns for table multi_stay.
var multiStayColumns = MultiStayColumns{
	Id:                   "id",
	CustomerName:         "customer_name",
	RoomNumber:           "room_number",
	PayTime:              "pay_time",
	CheckinTime:          "checkin_time",
	NumberOfNights:       "number_of_nights",
	Done:                 "done",
	BillId:               "bill_id",
	ExpectedCheckoutTime: "expected_checkout_time",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
	DeletedAt:            "deleted_at",
}

// NewMultiStayDao creates and returns a new DAO object for table data access.
func NewMultiStayDao() *MultiStayDao {
	return &MultiStayDao{
		group:   "default",
		table:   "multi_stay",
		columns: multiStayColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MultiStayDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MultiStayDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MultiStayDao) Columns() MultiStayColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MultiStayDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MultiStayDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MultiStayDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
