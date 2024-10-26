// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// ReportLogs is the golang structure for table report_logs.
type ReportLogs struct {
	Id          int64     `json:"id"          orm:"id"           description:""` //
	OriginalUrl string    `json:"originalUrl" orm:"original_url" description:""` //
	Method      string    `json:"method"      orm:"method"       description:""` //
	RequestBody string    `json:"requestBody" orm:"request_body" description:""` //
	CreatedAt   time.Time `json:"createdAt"   orm:"created_at"   description:""` //
	UpdatedAt   time.Time `json:"updatedAt"   orm:"updated_at"   description:""` //
	DeletedAt   time.Time `json:"deletedAt"   orm:"deleted_at"   description:""` //
}
