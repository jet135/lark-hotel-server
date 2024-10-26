// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"time"
)

// LarkAppInfo is the golang structure for table lark_app_info.
type LarkAppInfo struct {
	Id          int64     `json:"id"          orm:"id"           description:""` //
	AppName     string    `json:"appName"     orm:"app_name"     description:""` //
	AppToken    string    `json:"appToken"    orm:"app_token"    description:""` //
	FolderToken string    `json:"folderToken" orm:"folder_token" description:""` //
	Url         string    `json:"url"         orm:"url"          description:""` //
	BelongDate  time.Time `json:"belongDate"  orm:"belong_date"  description:""` //
	TableId     string    `json:"tableId"     orm:"table_id"     description:""` //
	DataSync    int       `json:"dataSync"    orm:"data_sync"    description:""` //
	CreatedAt   time.Time `json:"createdAt"   orm:"created_at"   description:""` //
	UpdatedAt   time.Time `json:"updatedAt"   orm:"updated_at"   description:""` //
	DeletedAt   time.Time `json:"deletedAt"   orm:"deleted_at"   description:""` //
}
