// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// Bill is the golang structure for table bill.
type Bill struct {
	Id                   string    `json:"id"                   orm:"id"                     description:""`                                                                             //
	Name                 string    `json:"name"                 orm:"name"                   description:"客户姓名"`                                                                         // 客户姓名
	RoomNumber           string    `json:"roomNumber"           orm:"room_number"            description:""`                                                                             //
	Phone                string    `json:"phone"                orm:"phone"                  description:""`                                                                             //
	Source               int       `json:"source"               orm:"source"                 description:""`                                                                             //
	SourceId             string    `json:"sourceId"             orm:"source_id"              description:""`                                                                             //
	RoomPrice            float64   `json:"roomPrice"            orm:"room_price"             description:""`                                                                             //
	RoomPaymentType      string    `json:"roomPaymentType"      orm:"room_payment_type"      description:""`                                                                             //
	Deposit              float64   `json:"deposit"              orm:"deposit"                description:""`                                                                             //
	DepositType          string    `json:"depositType"          orm:"deposit_type"           description:""`                                                                             //
	Amount               float64   `json:"amount"               orm:"amount"                 description:""`                                                                             //
	PayTime              time.Time `json:"payTime"              orm:"pay_time"               description:""`                                                                             //
	Shift                string    `json:"shift"                orm:"shift"                  description:""`                                                                             //
	Remark               string    `json:"remark"               orm:"remark"                 description:""`                                                                             //
	NumberOfNights       int       `json:"numberOfNights"       orm:"number_of_nights"       description:""`                                                                             //
	ExpectedCheckoutTime time.Time `json:"expectedCheckoutTime" orm:"expected_checkout_time" description:""`                                                                             //
	CheckinTime          time.Time `json:"checkinTime"          orm:"checkin_time"           description:""`                                                                             //
	CheckoutTime         time.Time `json:"checkoutTime"         orm:"checkout_time"          description:""`                                                                             //
	DepositRefundTime    time.Time `json:"depositRefundTime"    orm:"deposit_refund_time"    description:""`                                                                             //
	BelongDate           time.Time `json:"belongDate"           orm:"belong_date"            description:""`                                                                             //
	AssociatedBillId     string    `json:"associatedBillId"     orm:"associated_bill_id"     description:""`                                                                             //
	IsAdditionalItem     int       `json:"isAdditionalItem"     orm:"is_additional_item"     description:""`                                                                             //
	CreatedAt            time.Time `json:"createdAt"            orm:"created_at"             description:""`                                                                             //
	UpdatedAt            time.Time `json:"updatedAt"            orm:"updated_at"             description:""`                                                                             //
	DeletedAt            time.Time `json:"deletedAt"            orm:"deleted_at"             description:"Indicates logical deletion. When empty, it indicates that the data is valid."` // Indicates logical deletion. When empty, it indicates that the data is valid.
}
