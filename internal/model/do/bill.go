// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Bill is the golang structure of table bill for DAO operations like Where/Data.
type Bill struct {
	g.Meta               `orm:"table:bill, do:true"`
	Id                   interface{} //
	Name                 interface{} // 客户姓名
	RoomNumber           interface{} //
	Phone                interface{} //
	Source               interface{} //
	SourceId             interface{} //
	RoomPrice            interface{} //
	RoomPaymentType      interface{} //
	Deposit              interface{} //
	DepositType          interface{} //
	Amount               interface{} //
	PayTime              interface{} //
	Shift                interface{} //
	Remark               interface{} //
	NumberOfNights       interface{} //
	ExpectedCheckoutTime interface{} //
	CheckinTime          interface{} //
	CheckoutTime         interface{} //
	DepositRefundTime    interface{} //
	BelongDate           interface{} //
	AssociatedBillId     interface{} //
	IsAdditionalItem     interface{} //
	CreatedAt            interface{} //
	UpdatedAt            interface{} //
	DeletedAt            interface{} // Indicates logical deletion. When empty, it indicates that the data is valid.
}
