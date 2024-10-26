// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ReportLogs is the golang structure of table report_logs for DAO operations like Where/Data.
type ReportLogs struct {
	g.Meta      `orm:"table:report_logs, do:true"`
	Id          interface{} //
	OriginalUrl interface{} //
	Method      interface{} //
	RequestBody interface{} //
	CreatedAt   interface{} //
	UpdatedAt   interface{} //
	DeletedAt   interface{} //
}
