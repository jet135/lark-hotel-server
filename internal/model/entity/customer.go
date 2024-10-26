// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// Customer is the golang structure for table customer.
type Customer struct {
	Id               string    `json:"id"               orm:"id"                 description:""` //
	Name             string    `json:"name"             orm:"name"               description:""` //
	BirthDate        string    `json:"birthDate"        orm:"birth_date"         description:""` //
	Czdz             string    `json:"czdz"             orm:"czdz"               description:""` //
	IdCode           string    `json:"idCode"           orm:"id_code"            description:""` //
	IdType           string    `json:"idType"           orm:"id_type"            description:""` //
	Phone            string    `json:"phone"            orm:"phone"              description:""` //
	XzqhTitle        string    `json:"xzqhTitle"        orm:"xzqhTitle"          description:""` //
	RoomPriceHistory string    `json:"roomPriceHistory" orm:"room_price_history" description:""` //
	Remark           string    `json:"remark"           orm:"remark"             description:""` //
	LastCheckInTime  time.Time `json:"lastCheckInTime"  orm:"last_check_in_time" description:""` //
	CreatedAt        time.Time `json:"createdAt"        orm:"created_at"         description:""` //
	UpdatedAt        time.Time `json:"updatedAt"        orm:"updated_at"         description:""` //
	DeletedAt        time.Time `json:"deletedAt"        orm:"deleted_at"         description:""` //
}
