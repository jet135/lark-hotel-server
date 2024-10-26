// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BillDao is the data access object for table bill.
type BillDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns BillColumns // columns contains all the column names of Table for convenient usage.
}

// BillColumns defines and stores column names for table bill.
type BillColumns struct {
	Id                   string //
	Name                 string // 客户姓名
	RoomNumber           string //
	Phone                string //
	Source               string //
	SourceId             string //
	RoomPrice            string //
	RoomPaymentType      string //
	Deposit              string //
	DepositType          string //
	Amount               string //
	PayTime              string //
	Shift                string //
	Remark               string //
	NumberOfNights       string //
	ExpectedCheckoutTime string //
	CheckinTime          string //
	CheckoutTime         string //
	DepositRefundTime    string //
	BelongDate           string //
	AssociatedBillId     string //
	IsAdditionalItem     string //
	CreatedAt            string //
	UpdatedAt            string //
	DeletedAt            string // Indicates logical deletion. When empty, it indicates that the data is valid.
}

// billColumns holds the columns for table bill.
var billColumns = BillColumns{
	Id:                   "id",
	Name:                 "name",
	RoomNumber:           "room_number",
	Phone:                "phone",
	Source:               "source",
	SourceId:             "source_id",
	RoomPrice:            "room_price",
	RoomPaymentType:      "room_payment_type",
	Deposit:              "deposit",
	DepositType:          "deposit_type",
	Amount:               "amount",
	PayTime:              "pay_time",
	Shift:                "shift",
	Remark:               "remark",
	NumberOfNights:       "number_of_nights",
	ExpectedCheckoutTime: "expected_checkout_time",
	CheckinTime:          "checkin_time",
	CheckoutTime:         "checkout_time",
	DepositRefundTime:    "deposit_refund_time",
	BelongDate:           "belong_date",
	AssociatedBillId:     "associated_bill_id",
	IsAdditionalItem:     "is_additional_item",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
	DeletedAt:            "deleted_at",
}

// NewBillDao creates and returns a new DAO object for table data access.
func NewBillDao() *BillDao {
	return &BillDao{
		group:   "default",
		table:   "bill",
		columns: billColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BillDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BillDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BillDao) Columns() BillColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BillDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BillDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BillDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
