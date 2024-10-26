// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// LarkAppInfo is the golang structure of table lark_app_info for DAO operations like Where/Data.
type LarkAppInfo struct {
	g.Meta      `orm:"table:lark_app_info, do:true"`
	Id          interface{} //
	AppName     interface{} //
	AppToken    interface{} //
	FolderToken interface{} //
	Url         interface{} //
	BelongDate  interface{} //
	TableId     interface{} //
	DataSync    interface{} //
	CreatedAt   interface{} //
	UpdatedAt   interface{} //
	DeletedAt   interface{} //
}
