// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MultiStay is the golang structure of table multi_stay for DAO operations like Where/Data.
type MultiStay struct {
	g.Meta               `orm:"table:multi_stay, do:true"`
	Id                   interface{} //
	CustomerName         interface{} //
	RoomNumber           interface{} //
	PayTime              interface{} //
	CheckinTime          interface{} //
	NumberOfNights       interface{} //
	Done                 interface{} //
	BillId               interface{} //
	ExpectedCheckoutTime interface{} //
	CreatedAt            interface{} //
	UpdatedAt            interface{} //
	DeletedAt            interface{} //
}
