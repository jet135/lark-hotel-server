// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// MultiStay is the golang structure for table multi_stay.
type MultiStay struct {
	Id                   int64     `json:"id"                   orm:"id"                     description:""` //
	CustomerName         string    `json:"customerName"         orm:"customer_name"          description:""` //
	RoomNumber           string    `json:"roomNumber"           orm:"room_number"            description:""` //
	PayTime              time.Time `json:"payTime"              orm:"pay_time"               description:""` //
	CheckinTime          time.Time `json:"checkinTime"          orm:"checkin_time"           description:""` //
	NumberOfNights       int       `json:"numberOfNights"       orm:"number_of_nights"       description:""` //
	Done                 int       `json:"done"                 orm:"done"                   description:""` //
	BillId               string    `json:"billId"               orm:"bill_id"                description:""` //
	ExpectedCheckoutTime time.Time `json:"expectedCheckoutTime" orm:"expected_checkout_time" description:""` //
	CreatedAt            time.Time `json:"createdAt"            orm:"created_at"             description:""` //
	UpdatedAt            time.Time `json:"updatedAt"            orm:"updated_at"             description:""` //
	DeletedAt            time.Time `json:"deletedAt"            orm:"deleted_at"             description:""` //
}
