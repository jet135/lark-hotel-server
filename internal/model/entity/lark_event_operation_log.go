// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// LarkEventOperationLog is the golang structure for table lark_event_operation_log.
type LarkEventOperationLog struct {
	Id        string    `json:"id"        orm:"id"         description:""` //
	MsgInfo   string    `json:"msgInfo"   orm:"msg_info"   description:""` //
	EventType string    `json:"eventType" orm:"event_type" description:""` //
	CreatedAt time.Time `json:"createdAt" orm:"created_at" description:""` //
	UpdatedAt time.Time `json:"updatedAt" orm:"updated_at" description:""` //
	DeletedAt time.Time `json:"deletedAt" orm:"deleted_at" description:""` //
}
